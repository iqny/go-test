package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	_interface "go-zh/ch32/queue/interface"
	_type "go-zh/ch32/queue/type"
)

var _ _interface.Derive = (*Kafka)(nil)

func New(config *Config) *Kafka {
	k := &Kafka{cfg: config}
	k.Connect()
	return k
}

type Kafka struct {
	group sarama.ConsumerGroup
	cfg   *Config
}

func (k *Kafka) Connect() error {
	var err error
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_11_0_2
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	k.group, err = sarama.NewConsumerGroup(k.cfg.Addr, k.cfg.GroupID, config)
	//defer k.group.Close()
	return err
}
func (k *Kafka) Ack(msg interface{}, b bool) {
	handle := msg.(*sessClaim)
	//fmt.Printf("Message topic:%q partition:%d offset:%d value:%s\n", handle.consumerMessage.Topic, handle.consumerMessage.Partition, handle.consumerMessage.Offset, string(handle.consumerMessage.Value))
	handle.sess.MarkMessage(handle.consumerMessage, "")
	handle.sess.Commit()
}

func (k *Kafka) Message(topics []string) <-chan *_type.Data {
	var consumerGroup *consumerGroupHandler
	consumerGroup = &consumerGroupHandler{name: "sera", sessClaim: make(chan *sessClaim, 1)}
	go func() {
		err := k.group.Consume(context.Background(), topics, consumerGroup)
		if err != nil {
			fmt.Println("t", err.Error())
		}
	}()
	data := make(chan *_type.Data)
	go func() {
		for consumer := range consumerGroup.sessClaim {
			data <- &_type.Data{Body: consumer.consumerMessage.Value, MessageHandel: consumer}
		}
	}()
	return data
}
func (k *Kafka) Close() error {
	return k.group.Close()
}

type consumerGroupHandler struct {
	name      string
	sessClaim chan *sessClaim
}
type sessClaim struct {
	sess            sarama.ConsumerGroupSession
	consumerMessage *sarama.ConsumerMessage
}

func (*consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	//i := 0
	for msg := range claim.Messages() {
		//fmt.Printf("Message topic:%q partition:%d offset:%d value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		/*sess.MarkMessage(msg, "")
		i++
		if i%15 == 0 {
		sess.Commit()
		}*/
		h.sessClaim <- &sessClaim{
			sess:            sess,
			consumerMessage: msg,
		}
	}
	return nil
}
