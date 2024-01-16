package app

import "github.com/rattapon001/porter-management/internal/job/domain"

func (s *JobUseCaseImpl) FindJobById(id domain.JobId) (*domain.Job, error) {
	job, err := s.Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return job, nil
}
