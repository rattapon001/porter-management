package app

import "github.com/rattapon001/porter-management/internal/job/domain"

type JobService interface {
	CreatedNewJob(location domain.Location, patient domain.Patient) (*domain.Job, error)
}

type JobServiceImpl struct {
	Repo domain.JobRepository
}

func NewJobService(repo domain.JobRepository) JobService {
	return &JobServiceImpl{
		Repo: repo,
	}
}
