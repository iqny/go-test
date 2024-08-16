package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个有缓存的channel
	ch := make(chan int, 3) //容量是3

	//len(ch)缓冲区剩余数据个数， cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch)= %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 10; i++ { //这里数据量大于管道容量，会出阻塞
			ch <- i //往chan写内容，如果主协程没读的话，写满3个就会阻塞在此
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch)= %d\n", i, len(ch), cap(ch))
		}
		close(ch)
	}()
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	//延时
	time.Sleep(2 * time.Second)

	for {
		num, ok := <-ch //读管道中内容，没有内容前，阻塞
		if ok {
			fmt.Println("num = ", num)
		} else {
			break
		}
	}
	fmt.Println("end")
}
