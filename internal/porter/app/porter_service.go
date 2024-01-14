package app

import (
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterService interface {
	CreatedNewPorter(name string) (*domain.Porter, error)
}

type PorterServiceImpl struct {
	Repo      domain.PorterRepository
	Publisher event.EventHandler
}

func NewPorterService(repo domain.PorterRepository, publisher event.EventHandler) PorterService {
	return &PorterServiceImpl{
		Repo:      repo,
		Publisher: publisher,
	}
}
