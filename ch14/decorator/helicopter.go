package decorator

import "fmt"

type Helicopter struct {}

func (h *Helicopter) fly() {
	fmt.Println("fly")
}

func (h *Helicopter) landing() {
	fmt.Println("landing")
}


