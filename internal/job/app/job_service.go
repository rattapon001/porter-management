package app

import (
	"github.com/rattapon001/porter-management/internal/job/app/command"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type JobService interface {
	CreatedNewJob(location domain.Location, patient domain.Patient, publisher command.EventHandler) (*domain.Job, error)
}

type JobServiceImpl struct {
	Repo domain.JobRepository
}

func NewJobService(repo domain.JobRepository) JobService {
	return &JobServiceImpl{
		Repo: repo,
	}
}
