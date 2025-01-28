package config

import (
	"WB2/Infrastucture/kafka/Services"
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct {
	Brokers []string `json:"brokers,omitempty"`
	Topic   string   `json:"topic,omitempty"`
	GroupId string   `json:"group_id,omitempty"`
}

func InitKafkaConfig() *KafkaConfig {
	kafkaConfig := &KafkaConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "TestTopic",
		GroupId: "my-consumer-group",
	}
	return kafkaConfig
}

func (kc *KafkaConfig) NewKafkaConsumer() *Services.KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  kc.Brokers,
		GroupID:  kc.GroupId,
		Topic:    kc.Topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &Services.KafkaConsumer{Reader: reader}
}

func (kc *KafkaConfig) NewKafkaProducer() *Services.KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kc.Brokers...),
		Topic:    kc.Topic,
		Balancer: &kafka.RoundRobin{},
	}
	return &Services.KafkaProducer{Writer: writer}
}
