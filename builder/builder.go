package main

import "fmt"

type ResourceParams struct {
	name     string
	maxTotal int64
	maxIdle  int64
	minIdle  int64
}
type ResourceProduct interface {
	show()
}
type RedisResourceProduct struct {
	resourceParams ResourceParams
}

func (p *RedisResourceProduct) show() {
	fmt.Printf("Product的数据为 %+v ", p.resourceParams)
}
func main() {

}
