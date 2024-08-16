package main

import (
	"context"
	"fmt"
	job "go-zh/ch32/jobs"
	"go-zh/ch32/queue"
	"go-zh/ch32/queue/config"
	"go-zh/ch32/queue/drivers/kafka"
	"go-zh/ch32/queue/drivers/rabbitmq"
	_type "go-zh/ch32/queue/type"
)

func main() {
	/*var drvies _interface.Derive

	drvies = kafka.New(&kafka.Config{
		Addr:    []string{"172.16.14.153:9092"},
		GroupID: "t1",
	})*/
	/*for msg := range drvies.Message([]string{"web_log"}) {
		fmt.Printf("%s\n", msg.Body)
		drvies.Ack(msg.MessageHandel, true)
	}*/

	/*drvies = rabbitmq.New()
	for msg := range drvies.Message() {
		fmt.Printf("%s\n", msg.Body)
		drvies.Ack(msg.MessageHandel, true)
	}*/

	queueHandle := queue.NewJob(context.Background(), &config.Config{
		Driver:   "kafka",
		LogPath:  "",
		Exchange: "",
		Queue: []config.QueueConfig{{
			Name:        "test",
			Enable:      true,
			Count:       2,
			MaxRetryNum: 3,
		}},
		DeadName: "",
		KafkaConfig: &kafka.Config{
			Addr:    []string{"172.16.14.153:9092"},
			GroupID: "t1",
		},
		RabbitmqConfig: &rabbitmq.Config{
			Addr:     "amqp://admin:admin@127.0.0.1:5672/",
			Exchange: "my_exchange",
		},
	})
	fmt.Printf("%#v\n", job.Jobs)
	queueHandle.Register(job.Jobs)
	queueHandle.Run()
}

type mint struct {
	jobs map[string]_type.HandlerFunc
	ctx  context.Context
}

func run() {

}
