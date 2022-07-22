package main

import "testing"

func TestName(t *testing.T) {

}

func BenchmarkName(b *testing.B) {
	var auto = &AutoId{1005: &data{
		id:   10000000,
		incr: 1000,
		num:  1000,
	}}
	for i := 0; i < b.N; i++ {
			auto.GetId(1005)
	}
}
