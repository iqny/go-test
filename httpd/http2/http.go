package main

import (
	"fmt"
	"net/http"
)
type MyMux struct {}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,http"))
}

func main() {
	serve:=http.NewServeMux()
	serve.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"%s","hello,world.")
	})
	serve.Handle("/http",new(MyMux))
	http.ListenAndServe(":8080",serve)
}
