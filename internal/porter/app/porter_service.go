package app

import (
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterService interface {
	CreatedNewPorter(name string) (*domain.Porter, error)
	ReadyForJob(porter *domain.Porter, token string) error
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
