package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6}
	sum(i...)
}
func sum(i ...int) {
	for _, v := range i {
		fmt.Println(v)
	}
}
