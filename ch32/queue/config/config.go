package config

import (
	"go-zh/ch32/queue/drivers/kafka"
	"go-zh/ch32/queue/drivers/rabbitmq"
)

type Config struct {
	Driver         string
	LogPath        string
	Exchange       string
	Queue          []QueueConfig
	DeadName       string
	KafkaConfig    *kafka.Config
	RabbitmqConfig *rabbitmq.Config
}
type QueueConfig struct {
	Name        string
	Enable      bool
	Count       int64
	MaxRetryNum uint8
}
