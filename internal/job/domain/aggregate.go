package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rattapon001/porter-management/pkg"
)

const (
	JobCreatedEvent             pkg.EventName = "job_created"
	JobAcceptedEvent            pkg.EventName = "job_accepted"
	JobWorkingEvent             pkg.EventName = "job_working"
	JobCompletedEvent           pkg.EventName = "job_completed"
	JobAllocatedEvent           pkg.EventName = "job_allocated"
	JobEquipmentsNotEnoughEvent pkg.EventName = "job_equipments_not_enough"
	ItemAllocatedEvent          pkg.EventName = "item_allocated"
)

type Aggregate struct {
	Events []pkg.Event
}

func (a Aggregate) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Aggregate) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		return json.Unmarshal(data, a)
	}
	return fmt.Errorf("failed to unmarshal AggregateDB value: %v", value)
}

func (a *Aggregate) AppendEvent(id JobId, eventName pkg.EventName, payload interface{}) {
	a.Events = append(a.Events, pkg.Event{EventName: eventName, Payload: payload, EventTime: time.Now(), EventID: string(id)})
}
