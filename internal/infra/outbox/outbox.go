package outbox

import "github.com/google/uuid"

type OutboxEvent struct {
	ID            string      `json:"id"`
	AggregateType string      `json:"aggregatetype" gorm:"type:varchar(255);column:aggregatetype"`
	AggregateID   string      `json:"aggregateid" gorm:"type:varchar(255);column:aggregateid"`
	Type          string      `json:"type" gorm:"type:varchar(255);column:type"`
	Payload       interface{} `json:"payload" gorm:"type:jsonb;column:payload"`
}

func NewOutboxEvent(aggregateID, aggregateType, eventType string, payload interface{}) *OutboxEvent {
	return &OutboxEvent{
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		Type:          eventType,
		Payload:       payload,
		ID:            uuid.New().String(),
	}
}
