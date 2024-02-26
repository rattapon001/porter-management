package app

import (
	"context"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobUseCaseImpl) CreateNewJob(ctx context.Context, location domain.Location, patient domain.Patient, equipments []domain.Equipment) (*domain.Job, error) {

	var jobResult *domain.Job
	err := s.Uow.DoInTx(ctx, func(store uow.UnitOfWorkStore) error {
		job, err := domain.NewJob(location, patient, equipments)
		if err != nil {
			return err
		}
		err = store.Job().Save(job)
		if err != nil {
			return err
		}
		if err := s.Publisher.Publish(job.Aggregate.Events, store); err != nil {
			return err
		}
		jobResult = job

		return nil

	})
	if err != nil {
		return nil, err
	}

	return jobResult, nil
}
