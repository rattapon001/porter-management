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

func (r *JobMemoryRepository) Update(job *domain.Job) error {
	for i, j := range r.jobs {
		if j.ID == job.ID {
			if j.Version != job.Version {
				return infraErrors.ErrVersionMismatch
			}
			job.Version++
			r.jobs[i] = job
			return nil
		}
	}
	return nil
}
