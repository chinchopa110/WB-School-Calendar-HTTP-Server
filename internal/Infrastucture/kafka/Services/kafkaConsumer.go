package Services

import (
	"WB2/internal/Application/Domain"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

func (kc *KafkaConsumer) Close() error {
	return kc.Reader.Close()
}

func (kc *KafkaConsumer) Listen(ctx context.Context) *Domain.User {
	for {
		select {
		case <-ctx.Done():
			log.Println("consumer is shutdown")
			return nil
		default:
			m, err := kc.Reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("error while reading message: %v", err)
				time.Sleep(time.Second)
				continue
			}
			var user Domain.User
			if err := json.Unmarshal(m.Value, &user); err != nil {
				log.Printf("error unmarshalling message: %v", err)
				continue
			}
			return &user
		}
	}
}

func (kc *KafkaConsumer) processMessage(user *Domain.User) {
	log.Printf("Received user: %+v\n", user)
	for _, event := range user.Events {
		log.Printf("  - Event: %+v\n", event)
	}
}
