package event

type EventName string

type Event struct {
	EventName EventName
	Payload   interface{}
}

type EventHandler interface {
	Publish(event []Event) error
}
