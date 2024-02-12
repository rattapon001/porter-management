package kafka

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rattapon001/porter-management/internal/stock/app"
	"github.com/rattapon001/porter-management/internal/stock/app/command"
	"github.com/rattapon001/porter-management/pkg"
)

type kafkaConsumer struct {
	consumer     *kafka.Consumer
	StockUseCase app.StockUseCase
}

func NewKafkaConsumer(consumer *kafka.Consumer, useCase app.StockUseCase) *kafkaConsumer {
	return &kafkaConsumer{consumer: consumer, StockUseCase: useCase}
}

func (k *kafkaConsumer) Subscribe(topics []string) error {
	err := k.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	eventHandlers := map[string]command.StockCommand{
		"job_created": &command.StockAllocateCommand{},
	}
	go func() {
		run := true
		for run {
			ev := k.consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				_, err := k.consumer.CommitMessage(e)
				if err == nil {
					fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, e.Value)
					topics := e.TopicPartition.Topic
					if handler, ok := eventHandlers[*topics]; ok {
						handler.Execute(*topics, e.Value)
					} else {
						fmt.Println("No handler found for event : ", *topics)
					}
					// msg_process(e)
				}

			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				// fmt.Printf("Ignored %v\n", e)
			}
		}
	}()

	return nil
}

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(producer *kafka.Producer) *kafkaProducer {
	return &kafkaProducer{producer: producer}
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
