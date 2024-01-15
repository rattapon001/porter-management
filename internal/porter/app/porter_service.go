package app

import (
	"github.com/rattapon001/porter-management/internal/notification/app"
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterService interface {
	CreatedNewPorter(name string, token string) (*domain.Porter, error)
	ReadyForJob(ID domain.PorterId) error
	PorterAllowcated(payload domain.Job) error
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
