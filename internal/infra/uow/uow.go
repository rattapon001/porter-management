package uow

import (
	"context"
	"fmt"
	"time"

	infra_errors "github.com/rattapon001/porter-management/internal/infra/errors"
	outbox "github.com/rattapon001/porter-management/internal/infra/outbox"
	"github.com/rattapon001/porter-management/internal/job/domain"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	porter_domain "github.com/rattapon001/porter-management/internal/porter/domain"
	porter_postgres "github.com/rattapon001/porter-management/internal/porter/infra/postgres_orm"
	stock_domain "github.com/rattapon001/porter-management/internal/stock/domain"
	stock_postgres "github.com/rattapon001/porter-management/internal/stock/infra/postgres_orm"
	"gorm.io/gorm"
)

var errRetryCondition = map[string]bool{
	infra_errors.ErrVersionMismatch.Error(): true,
	gorm.ErrDuplicatedKey.Error():           true,
}

type uowStore struct {
	job    domain.JobRepository
	item   stock_domain.ItemRepository
	porter porter_domain.PorterRepository
	outbox outbox.OutboxRepository
}

type UnitOfWorkStore interface {
	Job() domain.JobRepository
	Item() stock_domain.ItemRepository
	Porter() porter_domain.PorterRepository
	Outbox() outbox.OutboxRepository
}

func (u *uowStore) Job() domain.JobRepository {
	return u.job
}

func (u *uowStore) Item() stock_domain.ItemRepository {
	return u.item
}

func (u *uowStore) Porter() porter_domain.PorterRepository {
	return u.porter
}

func (u *uowStore) Outbox() outbox.OutboxRepository {
	return u.outbox
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
			fmt.Printf("cannot retry error : %v\n", err)
			return err
		}

	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func (u *unitOfWork) DoInTx(ctx context.Context, fn UnitOfWorkBlock) (err error) {
	maxRetries := 3               // Maximum number of retries
	retryDelay := 2 * time.Second // Delay between retries

	return retry(maxRetries, retryDelay, func() error {
		return u.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			uowStore := &uowStore{
				job:    job_postgres.NewPostgresOrmRepository(tx),
				item:   stock_postgres.NewPostgresOrmRepository(tx),
				porter: porter_postgres.NewPostgresOrmRepository(tx),
				outbox: outbox.NewPostgresOutboxRepository(tx),
			}
			return fn(uowStore)
		})
	})

}
