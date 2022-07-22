package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}
type HelloService struct{}

func (h HelloService) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func main() {
	rpc.Register(new(HelloService))

	addr, err := net.ResolveTCPAddr("tcp", ":8090")
	if err != nil {
		log.Fatal("resolve:", err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal("listen:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}

}
