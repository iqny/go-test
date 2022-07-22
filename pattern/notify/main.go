package main

import "fmt"

type Notify interface {
	Send()
}
type Message struct {
}

func (m *Message) Send() { fmt.Println("implement me") }

type SNS struct {
}

func (S *SNS) Send()      { fmt.Println("implement me") }
func LetDo(notify Notify) { notify.Send() }

func main() {
	LetDo(new(Message))
	ScheduleUpdater(new(RedisConfig))
	ScheduleUpdater(new(KafKa))
	ServerShow(new(MysqlConfig))
	ServerShow(new(RedisConfig))
}

type Updater interface {
	Update() bool
}
type Shower interface {
	Show() string
}

type RedisConfig struct {
}

func (r *RedisConfig) Show() string {
	return "implement me"
}

func (r *RedisConfig) Update() bool {
	return true
}

type MysqlConfig struct {
}

func (m *MysqlConfig) Show() string {
	return "implement me"
}

type KafKa struct {
}

func (k *KafKa) Update() bool {
	return false
}
func ScheduleUpdater(updater Updater) bool {
	return updater.Update()
}
func ServerShow(shower Shower) string {
	return shower.Show()
}
