package nonempty

import (
	"fmt"
	"unicode"
)

func nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

// remove 2,3,4,5,6
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
func reverse(arr *[6]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func rotate(s []int, r int) []int {
	lens := len(s)
	arr := make([]int, lens)
	for k := range s {
		index := r + k
		if index >= lens {
			index -= lens
		}
		arr[k] = s[index]
	}
	return arr
}
func rotate2(t []int, n int) []int {
	/*
		这个地方我起初写成
		tmp := t[:n]
		测试是发现始终输出：[3,4,5,3,4]
		才想起来这是 浅拷贝，是浮动的
		正确使用要使用深拷贝去给tmp赋值
	*/
	tmp := make([]int, n)
	copy(tmp, t[:n])
	copy(t[:len(t)-n], t[n:])
	copy(t[len(t)-n:], tmp[:])
	return t
}
func rotate3(t []int, n int) []int {
	//fmt.Printf("origin=%#x\n",unsafe.Pointer(&t))
	for ; n > 0; n-- {
		v := t[0]
		copy(t[0:], t[1:])
		t[len(t)-1] = v
		//fmt.Printf("t=%#x\n",unsafe.Pointer(&t))
	}
	return t
}

// unique ["a","b","b","b","c"]
func unique(s []string) []string {
	k := 0
	for _, v := range s {
		if s[k] == v {
			continue
		}
		k++
		s[k] = v
	}
	return s[:k+1]
}

// emptyString []byte("abc bcd wer  sdsd  taoshihan     de")
func emptyString(b []byte) []byte {
	k := 0
	bLen := len(b)
	for i := 1; i < bLen-1; i++ {
		v := b[i]
		if v == b[k] && unicode.IsSpace(rune(v)) {
			continue
		}
		k++
		b[k] = v
	}
	return b[:k+1]
}
func reverseInPlace(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b[:]
}
func testMap() {
	m := make(map[string]int)

	if v, ok := m["key"]; ok {
		fmt.Println(v)
	}
}
func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
