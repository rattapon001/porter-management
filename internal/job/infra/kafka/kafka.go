package kafka

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/app/command"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type DLQsProducer interface {
	SendToDLQs(msg kafka.Message, processingErr error) error
}

type kafkaConsumer struct {
	consumer    *kafka.Consumer
	JobsUseCase app.JobUseCase
	dlq         DLQsProducer
}

type EventValue struct {
	Schema  interface{} `json:"schema"`
	Payload string      `json:"payload"`
}

func NewKafkaConsumer(consumer *kafka.Consumer, useCase app.JobUseCase, dlq DLQsProducer) *kafkaConsumer {
	return &kafkaConsumer{consumer: consumer, JobsUseCase: useCase, dlq: dlq}
}

func (k *kafkaConsumer) HandlerMessage(msg *kafka.Message) error {

	eventHandlers := map[string]command.JobCommand{
		string("job.events." + domain.ItemAllocatedEvent): &command.ItemAllocateCommand{
			JobsUseCase: k.JobsUseCase,
		},
	}

	fmt.Printf("Message on %s:\n%s\n", msg.TopicPartition, msg.Value)
	topic := msg.TopicPartition.Topic
	if handler, ok := eventHandlers[*topic]; ok {
		var data EventValue
		err := json.Unmarshal([]byte(msg.Value), &data)
		if err != nil {
			err := k.dlq.SendToDLQs(*msg, err)
			if err != nil {
				return err
			}
		}
		if err := handler.Execute(*topic, []byte(data.Payload)); err != nil {
			err := k.dlq.SendToDLQs(*msg, err)
			if err != nil {
				return err
			}
			return fmt.Errorf("error executing handler for event %s: %w", *topic, err)
		}
	} else {
		err := fmt.Errorf("no handler found for event: %s", *topic)
		err = k.dlq.SendToDLQs(*msg, err)
		if err != nil {
			return err
		}
		return err
	}

	return nil
}

func (k *kafkaConsumer) Subscribe(topics []string) error {

	err := k.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	go func() {
		run := true
		for run {
			ev := k.consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:

				if err := k.HandlerMessage(e); err != nil {
					fmt.Printf("Error handling message : %v\n", err)
				}

			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				continue
			}
		}
		k.consumer.Close()
	}()

	return nil
}
