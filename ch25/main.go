package main

import (
	"fmt"
	"sync"
	"time"
)

var quit = make(chan struct{})
var read = make(chan int)
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			read <- i
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("quit")
				return
			case i := <-read:
				fmt.Println("i:", i)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	quit <- struct{}{}
	wg.Wait()
}
