package postgresorm

import (
	"github.com/rattapon001/porter-management/internal/job/domain"
	"gorm.io/gorm"
)

type PostgresOrmRepository struct {
	db *gorm.DB
}

func NewPostgresOrmRepository(db *gorm.DB) *PostgresOrmRepository {
	return &PostgresOrmRepository{
		db: db,
	}
}

func (r *PostgresOrmRepository) Save(job *domain.Job) error {
	var jobDB domain.Job
	currentVersion := job.Version
	if err := r.db.First(&jobDB, job.ID).Error; err != nil {
		return r.db.Save(job).Error
	}
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
