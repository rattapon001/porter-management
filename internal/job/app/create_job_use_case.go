package app

import (
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobServiceImpl) CreatedNewJob(location domain.Location, patient domain.Patient) (*domain.Job, error) {
	job, err := domain.CreatedNewJob(location, patient)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Save(job)
	if err != nil {
		return nil, err
	}
	err = s.Publisher.Publish(job.Aggregate.Events)
	if err != nil {
		return nil, err
	}
	return job, nil
}
