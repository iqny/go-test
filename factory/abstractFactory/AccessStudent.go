package abstractFactory

import "fmt"

type AccessStudent struct{}

func (a *AccessStudent) Insert() bool {
	fmt.Println("AccessStudent Insert")
	return true
}

func (a *AccessStudent) Update() bool {
	fmt.Println("AccessStudent Update")
	return true
}
