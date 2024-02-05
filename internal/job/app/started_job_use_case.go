package app

import (
	"context"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobUseCaseImpl) StartJob(ctx context.Context, id domain.JobId) (*domain.Job, error) {

	var jobResult *domain.Job

	err := s.Uow.DoInTx(ctx, func(store uow.UnitOfWorkStore) error {

		job, err := s.Repo.FindById(id)
		if err != nil {
			return err
		}
		err = job.Start()
		if err != nil {
			return err
		}
		if err := store.Job().Save(job); err != nil {
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
