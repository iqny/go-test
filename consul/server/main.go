package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"go-zh/consul/api"
	"go-zh/consul/register"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)
type Order struct {
}

func (o *Order) List(ctx context.Context, in *api.PageRequest) (out *api.ListResponse, err error) {
	order := make([]*api.Order, 0)
	order = append(order,&api.Order{
		Id:                   1,
		Name:                 "1232123",
		Birthday: ptypes.TimestampNow(),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	order = append(order,&api.Order{
		Id:                   2,
		Name:                 "商品",
		Birthday: ptypes.TimestampNow(),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	order = append(order,&api.Order{
		Id:                   3,
		Name:                 "订单列表",
		Birthday: ptypes.TimestampNow(),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	out = &api.ListResponse{
		Order: order,
	}
	return
}

func main() {
	r:=NewRegister("192.168.99.101:8500","HelloService",8085)
	api.RegisterOrderServiceServer(r.Server,&Order{})
	r.Run()
}
//Registry 服务注册自定义结构体
type Registry struct {
	consulAddr,service string
	port int
	listener net.Listener
	Server *grpc.Server
	register *register.ConsulRegister
}
//NewRegister 创建新的服务注册
func NewRegister(consulAddr,service string,port int) *Registry {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v",port))
	if err != nil {
		log.Fatalln(err)
	}
	addrs := strings.Split(listener.Addr().String(),":")
	port,err = strconv.Atoi(addrs[len(addrs)-1])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("start server port :",addrs[len(addrs)-1])
	//consul service register
	nr := register.NewConsulRegister(consulAddr,service,port)
	nr.Register()
	//start grpc server
	serv :=  grpc.NewServer()
	//registe health check
	grpc_health_v1.RegisterHealthServer(serv, &register.HealthImpl{})

	return &Registry{consulAddr:consulAddr,service:service,port:port,listener:listener,Server:serv,register:nr}
}
//Run 启动
func (r *Registry)Run()  {
	//server hook
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
		log.Println("do run hook")
		r.register.Deregister()
		r.Server.Stop()
	}()

	if err := r.Server.Serve(r.listener); err != nil {
		panic(err)
	}
}

