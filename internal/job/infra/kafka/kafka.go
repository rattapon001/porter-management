package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rattapon001/porter-management/pkg"
)

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(kafka *kafka.Producer) *kafkaProducer {
	return &kafkaProducer{producer: kafka}
}

func (k *kafkaProducer) Publish(events []pkg.Event) error {

	event := events[len(events)-1]

	topic := string(event.EventName)
	message, err := json.Marshal(event.Payload)

	if err != nil {
		fmt.Printf("Error marshalling event: %v", err)
		return err
	}

	deliveryChan := make(chan kafka.Event)
	err = k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Error producing message: %v", err)
		return err
	}

	go func() {
		for e := range k.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	return nil
}
