package echo4

import (
	"fmt"
	"testing"
)

func TestEcho4(t *testing.T) {
	u:=[]uint8{'b'}
	fmt.Printf("%T\n",u)
	r:=rune(int32(1))
	fmt.Printf("%T\n",r)
	fmt.Println(-5%-3)
	var u1 uint8 = 255
	fmt.Printf("uint8 255=%08b\n",u1)
	fmt.Println(u1,u1+1,u1*u1)
	var i int8 = 127
	fmt.Println(i,i+1,i*i)
	fmt.Printf("127=%08b -128=%08b 1=%08b -2=%08b\n",i,i+1,i*i,-2)
	fmt.Printf("%d\n",0b11111111)
}
func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(100)
	}
}
