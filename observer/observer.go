package main

import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

type stduct struct {
	id   int64
	i322 int32
	name int
}

func main() {
	var s stduct
	fmt.Println(unsafe.Sizeof(s))
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	err := doFile("abd")
	fmt.Println(err)
	i := returnFunc()
	fmt.Println(i)
	var t1 T = T{"t1"}
	t1.M1()
	t1.M2()
	var t2 Intf = &t1
	t2.M1()
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func doFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		//return err
	}
	defer func(f *os.File) {
		err = errors.New("更新返回值")
	}(f)
	return
}
func returnFunc() (result int) {
	defer func() {
		p := recover()
		if p != nil {
			result = p.(int)
		}
	}()
	panic(13)
}

type T struct {
	Name string
}

func (t *T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}

type Intf interface {
	M1()
	M2()
}
