package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed message.txt
var message string

//go:embed static/result.png
var content []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		//fmt.Println(r.Header)
		fmt.Fprintf(w, "%s", message)
	})
	log.Fatal(http.ListenAndServe(":8989", nil))
}
