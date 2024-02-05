package uow

import (
	"context"
	"fmt"
	"time"

	infra_errors "github.com/rattapon001/porter-management/internal/infra/errors"
	"github.com/rattapon001/porter-management/internal/job/domain"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	"gorm.io/gorm"
)

var errRetryCondition = map[string]bool{
	infra_errors.ErrVersionMismatch.Error(): true,
	gorm.ErrDuplicatedKey.Error():           true,
}

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

// retry logic
func retry(attempts int, sleep time.Duration, fn func() error) (err error) {
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		if handler, ok := errRetryCondition[err.Error()]; ok && handler {
			fmt.Printf("Transaction failed: %v\n", err)
			time.Sleep(sleep)
			sleep *= 2
		} else {
			fmt.Printf("cannot retry: %v\n", err)
			return err
		}

	}
	fmt.Println("after loop")
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func (u *unitOfWork) DoInTx(ctx context.Context, fn UnitOfWorkBlock) (err error) {
	maxRetries := 3               // Maximum number of retries
	retryDelay := 2 * time.Second // Delay between retries

	return retry(maxRetries, retryDelay, func() error {
		return u.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			uowStore := &uowStore{
				jobs: job_postgres.NewPostgresOrmRepository(tx),
			}
			return fn(uowStore)
		})
	})

}
