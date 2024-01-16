package domain

import "github.com/rattapon001/porter-management/pkg"

const (
	JobCreatedEvent   pkg.EventName = "job_created"
	JobAcceptedEvent  pkg.EventName = "job_accepted"
	JobWorkingEvent   pkg.EventName = "job_working"
	JobCompletedEvent pkg.EventName = "job_completed"
)

type Aggregate struct {
	Events []pkg.Event
}

func (a *Aggregate) AppendEvent(eventName pkg.EventName, payload interface{}) {
	a.Events = append(a.Events, pkg.Event{EventName: eventName, Payload: payload})
}
