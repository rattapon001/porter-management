package app

import (
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterService interface {
	CreatedNewPorter(name string, token string) (*domain.Porter, error)
	ReadyForJob(ID domain.PorterId) error
}

type PorterServiceImpl struct {
	Repo      domain.PorterRepository
	Publisher event.EventHandler
	Pool      domain.Pool
}

func NewPorterService(repo domain.PorterRepository, publisher event.EventHandler, pool domain.Pool) PorterService {
	return &PorterServiceImpl{
		Repo:      repo,
		Publisher: publisher,
		Pool:      pool,
	}
}
