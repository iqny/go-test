package main

import "fmt"

func main() {

}

// 打印函数
func print[T any](s ...T) {
	for _, i := range s {
		fmt.Println(i)
	}
}
