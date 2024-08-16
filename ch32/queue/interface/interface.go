package _interface

import _type "go-zh/ch32/queue/type"

type Derive interface {
	Ack(msg interface{}, b bool)
	Connect() error
	Message(topics []string) <-chan *_type.Data
	Close() error
}
