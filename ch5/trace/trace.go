package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//bigSlowOperation()
	//double(12)
	fmt.Println(triple(4)) // "12"
	// fmt.Println(tr())
	// for _, f := range []string{"a", "b"} {
	// 	err := doFile(f)
	// 	fmt.Println(err)
	// }
}
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
func tr() (i int) {
	defer func() { fmt.Printf("tr()=%d\n", i) }()
	i = 2
	return 3
}
func doFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
