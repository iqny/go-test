package job

import (
	"go-zh/ch32/jobs/test"
	"go-zh/ch32/queue"
)

var Jobs = make(map[string]queue.Handler)

func init() {
	Register(test.New())
}

func Register(jobHandle queue.Handler) {
	Jobs[jobHandle.GetJobName()] = jobHandle
}
