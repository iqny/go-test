package main

import (
	"testing"
)

func BenchmarkStringJoin1(b *testing.B) {
	b.ReportAllocs()
	input := []string{"Hello", "World"}
	join:= func(strs []string,delim string)string {
		if len(strs) ==2{
			return strs[0]+delim+strs[1]
		}
		return ""
	}
	for i := 0; i < b.N; i++ {
		result := join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result:  " + result)
		}
	}
}
