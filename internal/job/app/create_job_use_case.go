package app

import (
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobServiceImpl) CreateNewJob(location domain.Location, patient domain.Patient) (*domain.Job, error) {
	job, err := domain.NewJob(location, patient)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Save(job)
	if err != nil {
		return nil, err
	}
	if err := s.Publisher.Publish(job.Aggregate.Events); err != nil {
		return nil, err
	}
	return job, nil
}
