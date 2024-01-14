package app

import (
	event "github.com/rattapon001/porter-management/internal/job/app/event_handler"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobServiceImpl) CreatedNewJob(location domain.Location, patient domain.Patient, publisher event.EventHandler) (*domain.Job, error) {
	job, err := domain.CreatedNewJob(location, patient)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Save(job)
	if err != nil {
		return nil, err
	}
	err = publisher.Publish(job.Aggregate.Events)
	if err != nil {
		return nil, err
	}
	return job, nil
}
