package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func NewKafkaProducer() (*kafka.Producer, error) {
	config := getKafkaConfig()
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func NewKafkaConsumer() (*kafka.Consumer, error) {
	config := GetKafkaConsumerConfig()
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
