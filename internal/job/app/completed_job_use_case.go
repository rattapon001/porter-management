package app

import "github.com/rattapon001/porter-management/internal/job/domain"

func (s *JobServiceImpl) CompletedJob(id domain.JobId) (*domain.Job, error) {
	job, err := s.Repo.FindById(domain.JobId(id))
	if err != nil {
		return nil, err
	}
	job.CompletedJob()
	if err := s.Repo.Update(job); err != nil {
		return nil, err
	}

	if err := s.Publisher.Publish(job.Aggregate.Events); err != nil {
		return nil, err
	}

	return job, nil
}
