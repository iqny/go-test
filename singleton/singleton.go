package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var (
	once   sync.Once
	single *singleton
)

func (s *singleton) Show() {
	fmt.Println("Show")
}
// GetSingleInstance 单例
func GetSingleInstance() *singleton {
	once.Do(func() {
		single = &singleton{}
	})
	return single
}
