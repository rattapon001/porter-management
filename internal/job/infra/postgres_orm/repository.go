package postgresorm

import (
	"github.com/rattapon001/porter-management/internal/job/domain"
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
// It first checks if the job already exists in the database by querying the job ID.
// If the job does not exist, it inserts the job into the database.
// If the job already exists, it updates the job's version and other fields.
// The job's version is incremented before updating the job in the database.
// The update is performed only if the current version in the database matches the version of the job being updated.
// Returns an error if any database operation fails.
func (r *PostgresOrmRepository) Save(job *domain.Job) error {
	var jobDB domain.Job
	currentVersion := job.Version
	if err := r.db.First(&jobDB, job.ID).Error; err != nil {
		return r.db.Save(job).Error
	}
	// Job already exists, check version for optimistic locking
	job.Version++
	return r.db.Model(&jobDB).Updates(job).Where("version = ?", currentVersion).Error
}

func (r *PostgresOrmRepository) FindById(id domain.JobId) (*domain.Job, error) {
	var job domain.Job
	if err := r.db.First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}
