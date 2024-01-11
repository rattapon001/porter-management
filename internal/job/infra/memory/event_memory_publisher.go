package memory

import (
	"fmt"

	"github.com/rattapon001/porter-management/internal/job/app/command"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type EventHandlerImpl struct {
	Publisher interface{}
}

func NewMemoryEventHandler() command.EventHandler {
	return &EventHandlerImpl{
		Publisher: nil,
	}
}

func (h *EventHandlerImpl) Publish(event []domain.Event) error {
	fmt.Printf("Publishing event: %v\n", event)
	return nil
}
