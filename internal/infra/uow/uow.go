package uow

import (
	"context"

	"github.com/rattapon001/porter-management/internal/job/domain"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	"gorm.io/gorm"
)

type uowStore struct {
	jobs domain.JobRepository
}

type UnitOfWorkStore interface {
	Jobs() domain.JobRepository
}

func (u *uowStore) Jobs() domain.JobRepository {
	return u.jobs
}

type UnitOfWorkBlock func(UnitOfWorkStore) error

type UnitOfWork interface {
	DoInTx(ctx context.Context, fn UnitOfWorkBlock) error
}

type unitOfWork struct {
	DB *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		DB: db,
	}
}

func (u *unitOfWork) DoInTx(ctx context.Context, fn UnitOfWorkBlock) error {
	return u.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		uowStore := &uowStore{
			jobs: job_postgres.NewPostgresOrmRepository(tx),
		}
		return fn(uowStore)
	})
}
