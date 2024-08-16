package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// 定义一个全局翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Son        []Son  `json:"son" binding:"gt=0,dive,required"`
}
type Son struct {
	UUID string  `json:"UUID"  binding:"required"`
	OP   []int64 `json:"OP" binding:"required"`
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			// 获取validator.ValidationErrors类型的errors
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// 非validator.ValidationErrors类型错误直接返回
				fmt.Println(err.Error())
				errSlice := strings.Split(err.Error(), ".")
				s := strings.Split(errSlice[1], " ")
				if len(s) < 4 {
					s = strings.Split(errSlice[2], " ")
				}
				typeErrors := make(map[string]string)
				switch s[3] {
				case "uint8", "int8", "int", "uint", "uint16", "int16", "uint32", "int32", "uint64", "int64", "byte":
					typeErrors[s[0]] = fmt.Sprintf("%s类型不合法，请输入整型", s[0])
					//err = errors.New(fmt.Sprintf("%s 类型不正确，请输入整型", s[0]))
				case "float32", "float64":
					typeErrors[s[0]] = fmt.Sprintf("%s类型不合法，请输入浮点型", s[0])
					//err = errors.New(fmt.Sprintf("%s 类型不正确，请输入浮点型", s[0]))
				case "[]uint8", "[]int8", "[]int", "[]uint", "[]uint16", "[]int16", "[]uint32", "[]int32", "[]uint64", "[]int64", "[]byte", "[]float32", "[]float64":
					//err = errors.New(fmt.Sprintf("%s 类型不正确，请输入数组", s[0]))
					typeErrors[s[0]] = fmt.Sprintf("%s类型不合法，请输入数组", s[0])
				case "string":
					typeErrors[s[0]] = fmt.Sprintf("%s类型不合法,请输入字符串", s[0])
				default:
					index := strings.Index(s[3], "[]")
					if index >= 0 {
						typeErrors[s[0]] = fmt.Sprintf("%s类型不合法，请输入数组", s[0])
					} else {
						//err = errors.New(fmt.Sprintf("%s 类型不正确", s[0]))
						typeErrors[s[0]] = fmt.Sprintf("%s类型不合法", s[0])
					}

				}
				c.JSON(http.StatusOK, gin.H{
					"msg": typeErrors,
				})
				return
			}
			// validator.ValidationErrors类型错误则进行翻译
			// 并使用removeTopStruct函数去除字段名中的结构体名称标识
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
		// 保存入库等具体业务逻辑代码...

		c.JSON(http.StatusOK, "success")
	})

	_ = r.Run(":8999")
}
