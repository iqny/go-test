package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime:=time.Now()
	s := strings.Join(os.Args[1:], " ")
	fmt.Printf("%d\n",time.Now().Sub(startTime))
	fmt.Printf("%s\n", s)

}
