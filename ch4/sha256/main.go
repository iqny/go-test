package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x:=sha256.Sum256([]byte("x"))
	X:=sha256.Sum256([]byte("X"))
	fmt.Printf("x|sha256=%x\nX|sha256=%x\n%t\n%T\n%d\n",x,X,x==X,x,x)
	fmt.Printf("%x\n",45)
}
