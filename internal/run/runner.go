package run

import (
	Repos "WB2/internal/Infrastucture/DataAccess/Repositories"
	"WB2/internal/config"
	"log"
)

func Run() {
	connStr := "user=postgres password=123 dbname=wb2 sslmode=disable"
	db := config.GetUpSQL(connStr)

	userEventsRepo := Repos.NewUserEventsRepo(db)
	config.GetUpServer(userEventsRepo)

	kafkaConfig := config.InitKafkaConfig()
	consumer := kafkaConfig.NewKafkaConsumer()
	producer := kafkaConfig.NewKafkaProducer()

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("Failed to close consumer %v", err)
		}
		if err := producer.Close(); err != nil {
			log.Printf("Failed to close producer %v", err)
		}
		if err := db.Close(); err != nil {
			log.Printf("Could not close database connection: %s\n", err)
		}

	}()

	log.Println("Done.")
}
