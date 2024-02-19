package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func InitKafkaConsumer() (*kafka.Consumer, error) {
	config := GetKafkaConsumerConfig()
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
