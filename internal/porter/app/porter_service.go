package app

import (
	"github.com/rattapon001/porter-management/internal/notification/app"
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterService interface {
	CreatedNewPorter(name string, token string) (*domain.Porter, error)
	PorterAllowcated(payload domain.Job) (*domain.Porter, error)
	PorterWorking(ID domain.PorterId) (*domain.Porter, error)
	PorterAvailable(ID domain.PorterId) (*domain.Porter, error)
	PorterUnavailable(ID domain.PorterId) (*domain.Porter, error)
}

type PorterServiceImpl struct {
	Repo      domain.PorterRepository
	Publisher event.EventHandler
	Noti      app.NotificationService
}

func NewPorterService(repo domain.PorterRepository, publisher event.EventHandler) PorterService {
	return &PorterServiceImpl{
		Repo:      repo,
		Publisher: publisher,
	}
}
