package Services

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func (kp *KafkaProducer) SendMessage(ctx context.Context, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Value: jsonValue,
	}
	err = kp.Writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("failed to write message: %v\n", err)
		return err
	}
	log.Printf("message send to topic %v\n", msg)
	return nil
}

func (kp *KafkaProducer) Close() error {
	return kp.Writer.Close()
}
