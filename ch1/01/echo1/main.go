package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%s\n", s)
	var i, j int
	i++
	j = i
	fmt.Println(j)
}
