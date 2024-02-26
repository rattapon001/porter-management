package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rattapon001/porter-management/pkg"
)

const (
	ItemCreatedEvent   = "item_created"   // ItemCreatedEvent is an event when a item is created
	ItemUpdatedEvent   = "item_updated"   // ItemUpdatedEvent is an event when a item is updated
	ItemDeletedEvent   = "item_deleted"   // ItemDeletedEvent is an event when a item is deleted
	ItemAllocatedEvent = "item_allocated" // ItemAllocatedEvent is an event when a item is Allocated ** for jobs
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

func (a *Aggregate) AppendEvent(eventName pkg.EventName, payload interface{}, id string) {
	a.Events = append(a.Events, pkg.Event{EventName: eventName, Payload: payload, EventTime: time.Now(), EventID: id})
}
