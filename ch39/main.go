package main

import (
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/say", say)
	mux.HandleFunc("/img", GenerateImg)                                         //生成验证码图片API
	mux.Handle("/verify/", captcha.Server(captcha.StdWidth, captcha.StdHeight)) //刷新验证码API
	serve := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	log.Fatal(serve.ListenAndServe())
}
func say(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "%s,%s", "say", name)
}

// GenerateImg 生成验证码图片名称
func GenerateImg(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	bytes, _ := json.Marshal(map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": d.CaptchaId})
	w.Write(bytes)
}
