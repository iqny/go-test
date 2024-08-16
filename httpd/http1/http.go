package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("X-Proxy"))
		fmt.Fprintf(w, "%s proxy:%s", "hello,world。", r.Header.Get("X-Proxy"))
	})
	http.HandleFunc("/say", say)
	http.HandleFunc("/tur", tur)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func say(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("say,hello"))
}
func tur(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tur,jur"))
}
