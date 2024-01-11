package command

import "github.com/rattapon001/porter-management/internal/job/domain"

type EventHandler interface {
	publish(event domain.Event) error
}

type EventHandlerImpl struct {
	Publisher interface{}
}

func NewEventHandler(publisher interface{}) EventHandler {
	return &EventHandlerImpl{
		Publisher: publisher,
	}
}

func (h *EventHandlerImpl) publish(event domain.Event) error {
	return nil
}
