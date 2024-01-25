package postgresorm

import (
	"github.com/rattapon001/porter-management/internal/job/domain"
	infra_errors "github.com/rattapon001/porter-management/internal/job/infra/errors"
	"gorm.io/gorm"
)

type PostgresOrmRepository struct {
	db *gorm.DB
}

// NewPostgresOrmRepository creates a new instance of the PostgresOrmRepository.
// It takes a *gorm.DB as a parameter and returns a pointer to the PostgresOrmRepository.
func NewPostgresOrmRepository(db *gorm.DB) *PostgresOrmRepository {
	return &PostgresOrmRepository{
		db: db,
	}
}

// Save saves the given job to the Postgres database.
// If the job already exists, it updates the existing record.
// If the job doesn't exist, it inserts a new record.
// It returns an error if there is any issue with the database operation.
func (r *PostgresOrmRepository) Save(job *domain.Job) error {
	var existingJob domain.Job
	currentVersion := job.Version
	err := r.db.Where("id = ?", job.ID).First(&existingJob).Error
	if err == gorm.ErrRecordNotFound {
		// Job doesn't exist, insert it
		err = r.db.Create(job).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {

		if existingJob.Version != currentVersion {
			return infra_errors.ErrVersionMismatch
		}
		job.Version++
		err = r.db.Save(job).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresOrmRepository) FindById(id domain.JobId) (*domain.Job, error) {
	var job domain.Job
	if err := r.db.Preload("Equipments").First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}
