package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "hello,world。")
	})
	http.HandleFunc("/say",say)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func say(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("say,hello"))
}
