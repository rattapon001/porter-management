package app

import "github.com/rattapon001/porter-management/internal/job/domain"

func (s *JobServiceImpl) StartedJob(id domain.JobId) (*domain.Job, error) {
	job, err := s.Repo.FindById(domain.JobId(id))
	if err != nil {
		return nil, err
	}
	job.StartedJob()
	err = s.Repo.Update(job)
	if err != nil {
		return nil, err
	}
	err = s.Publisher.Publish(job.Aggregate.Events)
	if err != nil {
		return nil, err
	}
	return job, nil
}
