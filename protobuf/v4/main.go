package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := RegisterRestServiceHandlerFromEndpoint(
		ctx, mux, "localhost:5000",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}
	go httpd()
	http.ListenAndServe(":8080", mux)
}

func httpd() {
	grpcServer := grpc.NewServer()
	RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	lis, _ := net.Listen("tcp", ":5000")
	grpcServer.Serve(lis)
}

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}
