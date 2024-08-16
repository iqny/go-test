package queue

import (
	"encoding/json"
	"fmt"
	"go-zh/ch32/queue/config"
	_interface "go-zh/ch32/queue/interface"
	_type "go-zh/ch32/queue/type"
	"sync/atomic"
)

type Worker struct {
	cfg        config.QueueConfig
	dataHandle chan _type.Data
	jobs       map[string]_type.HandlerFunc
}

func NewWorker(cfg config.QueueConfig, jobs map[string]_type.HandlerFunc) *Worker {
	worker := &Worker{
		cfg:  cfg,
		jobs: jobs,
	}
	return worker
}
func (w *Worker) Run(derive _interface.Derive) {
	var count int64 = 0
	for {
		if w.cfg.Count <= atomic.LoadInt64(&count) {
			continue
		}
		atomic.AddInt64(&count, 1)
		go func() {
			defer atomic.AddInt64(&count, -1)
			w.do(derive)
		}()
	}
}
func (w *Worker) do(derive _interface.Derive) {
	for msg := range derive.Message([]string{w.cfg.Name}) {
		//接收消息
		var body _type.Sender
		var err error
		//var traceId string
		var retryNum uint8 = 0
		var traceId string
		json.Unmarshal(msg.Body, &body)
		if f, ok := w.jobs[body.Job]; ok {
			for w.cfg.MaxRetryNum >= retryNum {
				func() {
					defer func() {
						if errRecover := recover(); errRecover != nil {
							err = errRecover.(error)
						}
					}()
					traceId, err = f(body.Msg)
				}()
				if err == nil {
					derive.Ack(msg.MessageHandel, true)
					fmt.Printf("ack:ok;traceId:%s\n", traceId)
					break
				} else {
					fmt.Printf("call func err:%s\n", err)
					if retryNum == 0 {
						fmt.Printf("ack:fail,将进行%d次重试\n", w.cfg.MaxRetryNum)
					} else {
						fmt.Printf("第 %d 次消费失败 ((retryNum <= %v) \n", retryNum, w.cfg.MaxRetryNum)
					}
					retryNum++
				}
			}
			//fmt.Println("traceId",traceId)
		}

	}
}
