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
	client, err := rpc.DialHTTP("tcp", ":8090")
	if err != nil {
		log.Fatal("dial:",err)
	}
	args:=Args{
		A: 10,
		B: 2,
	}
	var reply int

	err = client.Call("HelloService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("hello:",err)
	}
	fmt.Printf("HelloService: %d*%d=%d\n", args.A, args.B, reply)
}
