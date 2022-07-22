package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type HelloService struct{}
type Args struct {
	A, B int
}

func (h *HelloService) Multiply(args *Args, reply *int) error {
	*reply = args.B * args.A
	return nil
}
func main() {
		svc :=new(HelloService)
		rpc.Register(svc)
		rpc.HandleHTTP()
		log.Fatal(http.ListenAndServe(":8090",nil))
}
