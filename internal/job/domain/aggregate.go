package domain

import "github.com/rattapon001/porter-management/pkg"

const (
	JobCreated   pkg.EventName = "job_created"
	JobAccepted  pkg.EventName = "job_accepted"
	JobWorking   pkg.EventName = "job_working"
	JobCompleted pkg.EventName = "job_completed"
)

type Aggregate struct {
	Events []pkg.Event
}

func (a *Aggregate) AppendEvent(eventName pkg.EventName, payload interface{}) {
	a.Events = append(a.Events, pkg.Event{EventName: eventName, Payload: payload})
}
