package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type data struct {
	id   int64
	incr int64
	num  int64
}

func (d *data) getId() int64 {
	//d.id++
	atomic.AddInt64(&d.id, +1)
	//d.num--
	atomic.AddInt64(&d.num, -1)
	if atomic.LoadInt64(&d.num) < 0 {
		return 0
	}
	return d.id
}
func (d *data) setId(id int64) {
	atomic.StoreInt64(&d.id, id)
	atomic.StoreInt64(&d.num, d.incr)
}

type AutoId map[int]*data

var mu sync.Mutex
var log int64 = 1000

func (a AutoId) GetId(code int) int64 {
	var id int64
	mu.Lock()
	defer mu.Unlock()
	if dat, ok := a[code]; ok {
		for i := 0; i < 2; i++ {
			id = dat.getId()
			if id <= 0 {
				dat.setId(log)
				continue
			} else {
				break
			}
		}
	}
	return id
}
func main() {
	var auto = &AutoId{1005: &data{
		id:   100,
		incr: 5,
		num:  5,
	}}
	var wg sync.WaitGroup
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(auto.GetId(1005))

		}()
		//fmt.Println(auto.GetId(1005))
	}
	wg.Wait()
}
