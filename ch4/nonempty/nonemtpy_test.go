package nonempty

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

// BenchmarkRotate3-4   	41464173	        27.75 ns/op
// BenchmarkRotate3-4   	41950371	        27.75 ns/op
func BenchmarkRotate3(b *testing.B) {
	b.ResetTimer()
	a1 := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		rotate3(a1, 5)
	}
}

// BenchmarkRotate-4   	29249286	        42.28 ns/op
func BenchmarkRotate(b *testing.B) {
	b.ResetTimer()
	a1 := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		rotate(a1, 5)
	}
}

// BenchmarkRotate2-4   	28349044	        44.01 ns/op
func BenchmarkRotate2(b *testing.B) {
	b.ResetTimer()
	a1 := []int{0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		rotate2(a1, 5)
	}
}
func TestUnique(t *testing.T) {
	s := []string{"a", "b", "b", "b", "c"}
	s = unique(s)
	t.Log(s)
}
func TestEmptyString(t *testing.T) {
	b := []byte("abc bcd wer  sdsdd  taoshihan     de")
	fmt.Printf("%s\n", emptyString(b))
}
func TestReverseInPlace(t *testing.T) {
	s:=reverseInPlace([]byte("世办暗抗"))
	t.Log(s)
}
func TestEqualMap(t *testing.T) {
	//mapEqual(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Printf("%q",[]string{"a","b"})
	fmt.Println(utf8.UTFMax)
	now := time.Now()
	prMonth := now.AddDate(0, -1, 0)
	prYear := now.AddDate(-1, 0, 0)
	fmt.Println(prMonth)
	fmt.Println(prYear)
	fmt.Println(strings.Map(func(r rune) rune {
		return r+1
	},"HAL-9000"))
}
