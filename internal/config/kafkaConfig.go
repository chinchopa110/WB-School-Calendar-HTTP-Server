package config

import (
	Services2 "WB2/internal/Infrastucture/kafka/Services"
	"github.com/segmentio/kafka-go"
	"log"
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

func (kc *KafkaConfig) NewKafkaConsumer() *Services2.KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  kc.Brokers,
		GroupID:  kc.GroupId,
		Topic:    kc.Topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	consumer := &Services2.KafkaConsumer{Reader: reader}

	if consumer == nil {
		log.Fatalf("Failed to create consumer")
	}

	return consumer
}

func (kc *KafkaConfig) NewKafkaProducer() *Services2.KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kc.Brokers...),
		Topic:    kc.Topic,
		Balancer: &kafka.RoundRobin{},
	}

	producer := &Services2.KafkaProducer{Writer: writer}

	if producer == nil {
		log.Fatalf("Failed to create producer")
	}

	return producer
}
