package app

import (
	"context"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	event "github.com/rattapon001/porter-management/internal/job/app/event_handler"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type JobUseCase interface {
	CreateNewJob(ctx context.Context, location domain.Location, patient domain.Patient, equipments []domain.Equipment) (*domain.Job, error)
	AcceptJob(ctx context.Context, id domain.JobId, porter domain.Porter) (*domain.Job, error)
	FindJobById(id domain.JobId) (*domain.Job, error)
	StartJob(ctx context.Context, id domain.JobId) (*domain.Job, error)
	CompleteJob(ctx context.Context, id domain.JobId) (*domain.Job, error)
	JobAllocate(ctx context.Context, id domain.JobId) (*domain.Job, error)
}

type JobUseCaseImpl struct {
	Repo      domain.JobRepository
	Publisher event.EventHandler
	Uow       uow.UnitOfWork
}

func NewJobUseCase(repo domain.JobRepository, publisher event.EventHandler, uow uow.UnitOfWork) JobUseCase {
	return &JobUseCaseImpl{
		Repo:      repo,
		Publisher: publisher,
		Uow:       uow,
	}
}
