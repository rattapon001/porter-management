package app

import "github.com/rattapon001/porter-management/internal/job/domain"

type JobServicePort interface {
	CreateJob(location domain.Location, patient domain.Patient) (*domain.Job, error)
}

type JobService struct {
	Repo domain.JobRepository
}

func NewJobService(repo domain.JobRepository) *JobService {
	return &JobService{
		Repo: repo,
	}
}
