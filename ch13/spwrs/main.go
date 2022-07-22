package main

import (
	"fmt"
	"sync"
)
// 交替打印 数字和字母
func main() {
	letter, number := make(chan bool), make(chan bool)
	quit := make(chan struct{})
	var wait sync.WaitGroup

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				letter <- true
			case <-quit:
				return
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					quit <- struct{}{}
					return
				}
				fmt.Println(string(i))
				i++
				fmt.Println(string(i))
				i++
				number <- true
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}
