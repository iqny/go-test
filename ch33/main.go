package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type form struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard //把日志丢弃掉
	g := gin.Default()
	g.POST("/nacos/v1/auth/login", func(c *gin.Context) {
		fmt.Printf("%#v\n", c.Request.Header)
		var f form
		err := c.ShouldBind(&f)
		if err != nil {
			fmt.Printf("err:%s", err)
		}
		fmt.Println(f)
		c.JSON(http.StatusOK, gin.H{"accessToken": "123123123"})
	})
	g.GET("/nacos/v1/cs/configs", func(c *gin.Context) {
		fmt.Printf("%#v\n", c.Request.Header)
	})
	g.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world!")
	})
	g.Run(":8090")
}
func t() {
	text := "oooooo灰灰"
	key := []byte("eOYxV@C6@MkHpoqu")
	// 初始向量
	iv := []byte("EasI$k70o1HLuWO1")

	// 加密
	result := CBCEncrypter([]byte(text), key, iv)
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	// 解密
	base64String, _ := base64.StdEncoding.DecodeString("7XTwiACevcSaHYd+NNlSRQ==")
	result1 := []byte(base64String)
	fmt.Println(string(CBCDecrypter(result1, key, iv)))
	var ts = make([]string, 10)
	ts = append(ts, "ag")
	fmt.Println(ts)
}

/*
CBCEncrypter
CBC 加密
text 待加密的明文
key 秘钥
*/
func CBCEncrypter(text []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// 填充
	paddText := PKCS7Padding(text, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 加密
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	// 返回密文
	return result
}

/*
CBCDecrypter
// CBC 解密
// encrypter 待解密的密文
// key 秘钥
*/
func CBCDecrypter(encrypter []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(encrypter))
	blockMode.CryptBlocks(result, encrypter)
	// 去除填充
	result = UnPKCS7Padding(result)
	return result
}

/*
PKCS7Padding 填充模式
text：明文内容
blockSize：分组块大小
*/
func PKCS7Padding(text []byte, blockSize int) []byte {
	// 计算待填充的长度
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

// UnPKCS7Padding
// 去除 PKCS7Padding 填充的数据
// text 待去除填充数据的原文
func UnPKCS7Padding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}
