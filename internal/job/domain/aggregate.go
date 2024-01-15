package domain

import "github.com/rattapon001/porter-management/pkg"

const (
	JobEventCreated   pkg.EventName = "job_created"
	JobEventAccepted  pkg.EventName = "job_accepted"
	JobEventWorking   pkg.EventName = "job_working"
	JobEventCompleted pkg.EventName = "job_completed"
)

type Aggregate struct {
	Events []pkg.Event
}

func (a *Aggregate) AppendEvent(eventName pkg.EventName, payload interface{}) {
	a.Events = append(a.Events, pkg.Event{EventName: eventName, Payload: payload})
}
