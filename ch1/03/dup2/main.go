package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if 0 == len(files) {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				if err != nil {
					return
				}
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
				return
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fs:=strings.Split(line,"-/")
			fmt.Printf("%d\t%s\t%s\n", n,fs[0],fs[1])
		}
	}
}
func countLines(file *os.File, counts map[string]int) map[string]int {

	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()+"-/"+file.Name()]++
	}
	return counts
}
