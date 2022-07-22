package main

import (
	hello_pb "go-zh/protobuf/pbgo"
	"log"
	"net/http"
)

type EchoService struct{}

func (p *EchoService) Echo(request *hello_pb.String, reply *hello_pb.String) error {
	*reply = *request
	return nil
}

func main() {
	router := hello_pb.EchoServiceHandler(new(EchoService))
	log.Fatal(http.ListenAndServe(":8080", router))
}
