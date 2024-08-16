package test

import (
	"fmt"
	"go-zh/ch32/queue"
)

type test struct {
	name string
	job  queue.Job
}

var t queue.Handler = &test{
	name: "test_job",
}

func New() queue.Handler {
	return t
}

func (t *test) Call(args interface{}) (traceId string, err error) {
	fmt.Println(args.(string))
	fmt.Println("test完成")
	return
}

func (t *test) GetJobName() string {
	return t.name
}
