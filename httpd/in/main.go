package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type myStruct struct {
	Id   int `json:"id" validate:"required" label:"id2"`
	Name int `json:"name" validate:"required" label:"名称"`
}

func main() {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	mystruct := myStruct{}
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(mystruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
		}
	}

	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")*/
}
