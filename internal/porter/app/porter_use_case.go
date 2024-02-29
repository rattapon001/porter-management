package app

import (
	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/internal/notification/app"
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterUseCase interface {
	CreateNewPorter(name string, token string) (*domain.Porter, error)
	PorterAllocate(payload domain.Job) (*domain.Porter, error)
	PorterWorking(code domain.PorterCode) (*domain.Porter, error)
	PorterAvailable(code domain.PorterCode) (*domain.Porter, error)
	PorterUnavailable(code domain.PorterCode) (*domain.Porter, error)
}

type PorterUseCaseImpl struct {
	Repo      domain.PorterRepository
	Publisher event.EventHandler
	Noti      app.NotificationService
	uow       uow.UnitOfWork
}

func NewPorterUseCase(repo domain.PorterRepository, publisher event.EventHandler, noti app.NotificationService, uow uow.UnitOfWork) PorterUseCase {
	return &PorterUseCaseImpl{
		Repo:      repo,
		Publisher: publisher,
		uow:       uow,
		Noti:      noti,
	}
}
