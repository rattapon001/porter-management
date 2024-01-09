package domain

type EventName string

const (
	JobCreated   EventName = "job_created"
	JobAccepted  EventName = "job_accepted"
	JobWorking   EventName = "job_working"
	JobCompleted EventName = "job_completed"
)

type Event struct {
	EventName EventName
	Payload   interface{}
}

type Aggregate struct {
	Events []Event
}

func (a *Aggregate) AppendEvent(eventName EventName, payload interface{}) {
	a.Events = append(a.Events, Event{EventName: eventName, Payload: payload})
}
