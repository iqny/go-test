package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Xkcd struct {
	Title      string
	Img        string
	Transcript string
}

func main() {
	var urls []string
	for i := 0; i < 1000; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		urls = append(urls, url)
	}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	fmt.Println(<-ch)
}
func fetch(url string, ch chan<- string) {
	var result Xkcd
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	titles := strings.Split(result.Transcript, " ")
	for _, v := range titles {
		if v == os.Args[1] {
			ch <- result.Img
			break
		}
	}
}
