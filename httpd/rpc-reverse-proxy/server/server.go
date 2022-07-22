package main

import (
	hello "go-zh/protobuf/v1"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (h *HelloService) Hello(s *hello.String, s2 *hello.String) error {
	return nil
}

func (h *HelloService) Say(s string, reply *string) error {
	*reply = s + ",world."
	return nil
}

func main() {
	//s := rpc.NewServer()

	//hello.RegisterHelloService(s, new(HelloService))
	rpc.Register(new(HelloService))
	for {
		conn, _ := net.Dial("tcp", ":8090")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
