package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	dial, err := jsonrpc.Dial("tcp", ":8090")
	if err != nil {
		log.Fatal("dial:", err)
	}
	var reply int
	args := Args{
		A: 2,
		B: 9,
	}
	dial.Call("HelloService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
