package memory

import (
	"github.com/rattapon001/porter-management/internal/job/domain"
	infraErrors "github.com/rattapon001/porter-management/internal/job/infra/errors"
)

type JobMemoryRepository struct {
	jobs []*domain.Job
}

func NewJobMemoryRepository() *JobMemoryRepository {
	return &JobMemoryRepository{
		jobs: []*domain.Job{},
	}
}

func (r *JobMemoryRepository) Save(job *domain.Job) error {
	for i, existingJob := range r.jobs {
		if existingJob.ID == job.ID {
			if existingJob.Version != job.Version {
				return infraErrors.ErrVersionMismatch
			}
			job.Version++
			r.jobs[i] = job
			return nil
		}
	}
	r.jobs = append(r.jobs, job)
	return nil
}

func (r *JobMemoryRepository) FindById(id domain.JobId) (*domain.Job, error) {
	for _, job := range r.jobs {
		if job.ID == id {
			return job, nil
		}
	}
	return nil, nil
}
