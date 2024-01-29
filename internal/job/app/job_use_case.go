package app

import (
	event "github.com/rattapon001/porter-management/internal/job/app/event_handler"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type JobUseCase interface {
	CreateNewJob(location domain.Location, patient domain.Patient, equipments []domain.Equipment) (*domain.Job, error)
	AcceptJob(id domain.JobId, porter domain.Porter) (*domain.Job, error)
	FindJobById(id domain.JobId) (*domain.Job, error)
	StartJob(id domain.JobId) (*domain.Job, error)
	CompleteJob(id domain.JobId) (*domain.Job, error)
	JobAllocate(id domain.JobId, equipments []domain.Equipment) (*domain.Job, error)
}

type JobUseCaseImpl struct {
	Repo      domain.JobRepository
	Publisher event.EventHandler
}

func NewJobUseCase(repo domain.JobRepository, publisher event.EventHandler) JobUseCase {
	return &JobUseCaseImpl{
		Repo:      repo,
		Publisher: publisher,
	}
}
