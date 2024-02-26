package publisher

import (
	"github.com/rattapon001/porter-management/internal/infra/outbox"
	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/pkg"
)

type Publisher struct{}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(events []pkg.Event, uow uow.UnitOfWorkStore) error {
	event := events[len(events)-1]
	outbox := outbox.NewOutboxEvent(string(event.EventID), string(event.EventName), string(event.EventName), event.Payload)
	if err := uow.Outbox().Save(*outbox); err != nil {
		return err
	}
	return nil
}
