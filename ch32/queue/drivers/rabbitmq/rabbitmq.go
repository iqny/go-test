package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	_interface "go-zh/ch32/queue/interface"
	_type "go-zh/ch32/queue/type"
	"sync"
	"time"
)

var _ _interface.Derive = (*RabbitMq)(nil)

type RabbitMq struct {
	conn     *amqp.Connection
	delivery chan amqp.Delivery
	ch       *amqp.Channel
	cfg      *Config
	topic    string
	mux      sync.Mutex
}

func New(cfg *Config) *RabbitMq {
	r := &RabbitMq{delivery: make(chan amqp.Delivery), cfg: cfg}
	r.Connect()
	return r
}
func (r *RabbitMq) Connect() error {
	var err error
	r.conn, err = amqp.Dial("amqp://admin:admin@127.0.0.1:5672/")
	if err != nil {
		return err
	}
	return nil
}
func (r *RabbitMq) Ack(msg interface{}, b bool) {
	handle := msg.(amqp.Delivery)
	handle.Ack(!b)
}
func (r *RabbitMq) Message(topics []string) <-chan *_type.Data {
	r.topic = topics[0]
	data := make(chan *_type.Data, 1)
	r.Bind()
	go func() {
		for delivery := range r.delivery {
			data <- &_type.Data{
				Body:          delivery.Body,
				MessageHandel: delivery,
			}
		}
	}()
	return data
}
func (r *RabbitMq) Bind() {
	var err error
	r.ch, err = r.conn.Channel()
	err = r.ch.ExchangeDeclare(
		//交换机名称
		r.cfg.Exchange, //@todo
		//交换机类型 广播类型
		amqp.ExchangeDirect,
		//是否持久化
		true,
		//是否字段删除
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		//是否阻塞 true表示要等待服务器的响应
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Failed to ExchangeDeclare:%v", err)
	}
	argsQue := make(map[string]interface{})
	queue, err := r.ch.QueueDeclare(
		r.topic,
		true, //是否持久化
		false,
		false,
		false,
		argsQue,
	)
	if err != nil {
		fmt.Printf("Failed to declare a queue:%v", err)
		return
	}
	err = r.ch.QueueBind(
		queue.Name,
		//在pub/sub模式下，这里的key要为空
		queue.Name,
		"my_exchange",
		false,
		nil,
	)
	/*err = ch.Qos(
		//每次队列只消费一个消息 这个消息处理不完服务器不会发送第二个消息过来
		//当前消费者一次能接受的最大消息数量
		100,
		//服务器传递的最大容量
		0,
		//如果为true 对channel可用 false则只对当前队列可用
		true,
	)*/
	consume, err := r.ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	go func() {
		for msg := range consume {
			r.delivery <- msg
		}
	}()
}

// 监控重试连接
func (r *RabbitMq) watch() {
	go func() {
		for {
			if r.conn == nil || r.conn.IsClosed() {
				r.mux.Lock()
				err := r.Connect()
				r.mux.Unlock()
				if err != nil {
					time.Sleep(1 * time.Second)
				}
			} else {
				time.Sleep(1 * time.Second)
			}
			/*select {
			case <-r.ctx.Done():
				return
			default:

			}*/
		}
	}()
}
func (r *RabbitMq) Close() error {
	return r.ch.Close()
}
