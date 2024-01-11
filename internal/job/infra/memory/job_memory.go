package memory

import "github.com/rattapon001/porter-management/internal/job/domain"

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
