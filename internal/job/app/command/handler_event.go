package command

import "github.com/rattapon001/porter-management/internal/job/domain"

type EventHandler interface {
	Publish(event []domain.Event) error
}
