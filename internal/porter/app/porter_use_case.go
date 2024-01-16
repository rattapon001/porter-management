package app

import (
	"github.com/rattapon001/porter-management/internal/notification/app"
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterUseCase interface {
	CreateNewPorter(name string, token string) (*domain.Porter, error)
	PorterAllowcate(payload domain.Job) (*domain.Porter, error)
	PorterWorking(code domain.PorterCode) (*domain.Porter, error)
	PorterAvailable(code domain.PorterCode) (*domain.Porter, error)
	PorterUnavailable(code domain.PorterCode) (*domain.Porter, error)
}

type PorterUseCaseImpl struct {
	Repo      domain.PorterRepository
	Publisher event.EventHandler
	Noti      app.NotificationService
}

func NewPorterUseCase(repo domain.PorterRepository, publisher event.EventHandler) PorterUseCase {
	return &PorterUseCaseImpl{
		Repo:      repo,
		Publisher: publisher,
	}
}
