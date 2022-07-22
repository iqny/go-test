package main

import (
	"fmt"
)

var quit chan func()
func main() {
		quit = make(chan func())
		go func() {
			for i := 1; i <= 10; i++ {
				j :=i
				quit<- func() {
					fmt.Println("first",j)
				}
			}
			close(quit)
		}()
		for f:=range quit{
			f()
		}
}
