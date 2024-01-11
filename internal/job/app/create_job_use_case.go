package app

import (
	"github.com/rattapon001/porter-management/internal/job/app/command"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobServiceImpl) CreatedNewJob(location domain.Location, patient domain.Patient, publisher command.EventHandler) (*domain.Job, error) {
	job, err := domain.CreateNewJob(location, patient)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Save(job)
	if err != nil {
		return nil, err
	}
	err = publisher.Publish(job.Aggregate.Events)
	return job, nil
}
