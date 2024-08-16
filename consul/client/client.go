package main

import (
	"context"
	"fmt"
	"go-zh/consul/api"
	"go-zh/consul/resolver"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn := NewClietnConn("127.0.0.1:8500", "HelloService")
	defer conn.Close()
	c := api.NewOrderServiceClient(conn)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		//调用远端的方法
		r, err := c.List(ctx, &api.PageRequest{
			Page: 1,
		})
		if err != nil {
			log.Printf("could not greet: %v", err)

		} else {
			//log.Printf("order: %s", r.Order)
			/*//获取的是系统时区的时间戳
			pbTimestamp := ptypes.TimestampNow()
			//此方法默认 UTC 时区
			goTime, _ := ptypes.Timestamp(pbTimestamp)
			//设定为系统时区
			fmt.Println(goTime.Local())*/
			for _, v := range r.Order {
				goTime := v.Birthday.AsTime()
				//goTime, _ := ptypes.Timestamp(v.Birthday)
				//设定为系统时区
				log.Printf("goods: %d,%s,%s", v.Id, v.Name, goTime.Local())
			}
		}
		time.Sleep(time.Second)
	}
}
func NewClietnConn(consulAddr, serviceName string) *grpc.ClientConn {
	schema, err := resolver.GenerateAndRegisterConsulResolver(consulAddr, serviceName)
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}
	//建立连接
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", schema, serviceName), grpc.WithInsecure(), grpc.WithDisableServiceConfig())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn

}
