package pkg

type EventName string

type Event struct {
	EventName EventName
	Payload   interface{}
}
