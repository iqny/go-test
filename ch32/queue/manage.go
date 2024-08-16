package queue

import (
	"go-zh/ch32/queue/config"
	"go-zh/ch32/queue/drivers/kafka"
	"go-zh/ch32/queue/drivers/rabbitmq"
	_interface "go-zh/ch32/queue/interface"
	_type "go-zh/ch32/queue/type"
	"sync"
	"time"
)

type Manage struct {
	cfg    *config.Config
	derive _interface.Derive
	jobs   map[string]_type.HandlerFunc
}

func New(cfg *config.Config) *Manage {
	return &Manage{cfg: cfg, jobs: map[string]_type.HandlerFunc{}}
}

func (m *Manage) Run() {
	var derive _interface.Derive
	switch m.cfg.Driver {
	case "rabbitmq":
		derive = rabbitmq.New(m.cfg.RabbitmqConfig) //kafka.New(m.cfg.KafkaConfig)
	case "kafka":
		derive = kafka.New(m.cfg.KafkaConfig)
	}
	consumes := make(map[string]struct{})
	mux := sync.Mutex{}
	for {
		for _, queue := range m.cfg.Queue {
			queueName := queue.Name
			mux.Lock()
			if _, ok := consumes[queueName]; ok {
				//fmt.Printf("%s已启动\n", queueName)
				time.Sleep(time.Second)
				mux.Unlock()
				continue
			}
			mux.Unlock()
			go func(queueCfg config.QueueConfig) {
				defer func() {
					//退出或异常退出
					mux.Lock()
					delete(consumes, queueCfg.Name)
					mux.Unlock()
				}()
				NewWorker(queueCfg, m.jobs).Run(derive)
			}(queue)
			consumes[queueName] = struct{}{}
		}
	}
}

// Register 注册任务回调函数
func (m *Manage) Register(jobName string, handelFunc _type.HandlerFunc) {
	m.jobs[jobName] = handelFunc
}
