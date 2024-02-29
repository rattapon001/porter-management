package kafka

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rattapon001/porter-management/internal/stock/app"
	"github.com/rattapon001/porter-management/internal/stock/app/command"
)

type DLQsProducer interface {
	SendToDLQs(msg kafka.Message, processingErr error) error
}

type kafkaConsumer struct {
	consumer     *kafka.Consumer
	StockUseCase app.StockUseCase
	dlq          DLQsProducer
}

type EventValue struct {
	Schema  interface{} `json:"schema"`
	Payload string      `json:"payload"`
}

func NewKafkaConsumer(consumer *kafka.Consumer, useCase app.StockUseCase, dlq DLQsProducer) *kafkaConsumer {
	return &kafkaConsumer{consumer: consumer, StockUseCase: useCase, dlq: dlq}
}

func (k *kafkaConsumer) HandlerMessage(msg *kafka.Message) error {

	eventHandlers := map[string]command.StockCommand{
		"job.events.job_created": &command.StockAllocateCommand{
			StockUseCase: k.StockUseCase,
		},
	}

	if _, err := k.consumer.CommitMessage(msg); err != nil {
		return fmt.Errorf("failed to commit message: %w", err)
	}

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
