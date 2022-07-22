package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	src := strings.NewReader("gol\n")

	dst := os.Stdout

	buffer:=make([]byte,1)

	bytes, err := io.CopyBuffer(dst, src, buffer)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("The number of bytes are: %d\n",bytes)
	b:=isASCII("MN")
	fmt.Println(b)
}

func isASCII(s string) bool {
	for _, c := range s {
		if c >= 0x80 || c == 0x00 {
			return false
		}
	}
	return true
}
