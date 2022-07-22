package abstractFactory

import "fmt"

type AccessScore struct{}

func (a *AccessScore) Insert() bool {
	fmt.Println("AccessScore Insert")
	return true
}

func (a *AccessScore) List() map[string]string {
	fmt.Println("AccessScore List")
	return nil
}
