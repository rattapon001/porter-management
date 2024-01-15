package app

import (
	event "github.com/rattapon001/porter-management/internal/job/app/event_handler"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type JobService interface {
	CreatedNewJob(location domain.Location, patient domain.Patient) (*domain.Job, error)
	AcceptedJob(id domain.JobId, porter domain.Porter) (*domain.Job, error)
	FindJobById(id domain.JobId) (*domain.Job, error)
	StartedJob(id domain.JobId) (*domain.Job, error)
}

type JobServiceImpl struct {
	Repo      domain.JobRepository
	Publisher event.EventHandler
}

func NewJobService(repo domain.JobRepository, publisher event.EventHandler) JobService {
	return &JobServiceImpl{
		Repo:      repo,
		Publisher: publisher,
	}
}
