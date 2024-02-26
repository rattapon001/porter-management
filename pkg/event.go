package pkg

import "time"

type EventName string

type Event struct {
	EventID   string      `bson:"eventID" json:"eventID"`
	EventName EventName   `bson:"eventName" json:"eventName"`
	Payload   interface{} `bson:"payload" json:"payload"`
	EventTime time.Time   `bson:"eventTime" json:"eventTime"`
}
