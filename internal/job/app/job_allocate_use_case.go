package app

import "github.com/rattapon001/porter-management/internal/job/domain"

func (s *JobUseCaseImpl) JobAllocate(id domain.JobId, equipments []domain.Equipment) (*domain.Job, error) {
	job, err := s.Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	err = job.Allocate()
	if err != nil {
		return nil, err
	}
	if err := s.Repo.Save(job); err != nil {
		return nil, err
	}
	if err := s.Publisher.Publish(job.Aggregate.Events); err != nil {
		return nil, err
	}
	return job, nil
}
