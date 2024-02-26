package event

import (
	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/pkg"
)

type EventHandler interface {
	Publish(event []pkg.Event, uow uow.UnitOfWorkStore) error
}
