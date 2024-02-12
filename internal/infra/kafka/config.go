package kafka

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func getKafkaConfig() *kafka.ConfigMap {

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}

	config := &kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
	}
	return config
}
