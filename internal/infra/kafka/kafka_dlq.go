package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type dlqProducer struct {
	producer *kafka.Producer
}

func NewKafkaDLQs() (*dlqProducer, error) {
	config := GetKafkaProducerConfig()
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return &dlqProducer{producer: producer}, nil
}

func (k *dlqProducer) SendToDLQs(msg kafka.Message, processingErr error) error {
	dlqTopic := "dlq"

	headers := []kafka.Header{{Key: "ProcessingError", Value: []byte(processingErr.Error())}}
	dlqMsg := &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &dlqTopic, Partition: kafka.PartitionAny}, Value: msg.Value, Headers: headers}
	err := k.producer.Produce(dlqMsg, nil)
	if err != nil {
		fmt.Println("Failed to produce message to DLQ: ", err)
		return fmt.Errorf("failed to produce message to DLQ: %w", err)
	}

	return nil
}
