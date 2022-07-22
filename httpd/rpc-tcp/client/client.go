package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", ":8090")
	if err != nil {
		log.Fatal("dial:", err)
	}
	args := Args{
		A: 2,
		B: 9,
	}
	var reply int
	err = client.Call("HelloService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("call:",err)
	}
	fmt.Printf("HelloService: %d*%d=%d\n", args.A, args.B, reply)
}
