package postgresorm

import (
	"github.com/rattapon001/porter-management/internal/porter/domain"
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

// Save saves the given Porter entity to the Postgres database.
// If the Porter already exists in the database, it updates the existing record.
// If the Porter doesn't exist, it inserts a new record.
// Returns an error if any error occurs during the save operation.
func (r *PostgresOrmRepository) Save(p *domain.Porter) error {
	var existingPorter domain.Porter
	err := r.db.Where("id = ?", p.ID).First(&existingPorter).Error
	if err == gorm.ErrRecordNotFound {
		// Porter doesn't exist, insert it
		err = r.db.Create(p).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		// Some error occurred during the find operation
		return err
	} else {

		err = r.db.Save(p).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// FindAvailablePorter retrieves an available porter from the Postgres ORM repository.
// It queries the database for a porter with the status set to domain.PorterStatusAvailable.
// If a matching porter is found, it returns a pointer to the porter.
// If no porter is found, it returns nil.
func (r *PostgresOrmRepository) FindAvailablePorter() (*domain.Porter, error) {
	var porter domain.Porter
	err := r.db.Where("status = ?", domain.PorterStatusAvailable).First(&porter).Error
	if err != nil {
		return nil, err
	}
	return &porter, nil
}

// FindByID retrieves a porter from the database based on the given ID.
// It returns a pointer to the porter and an error if any occurred.
func (r *PostgresOrmRepository) FindByID(id domain.PorterId) (*domain.Porter, error) {
	var porter domain.Porter
	err := r.db.Where("id = ?", id).First(&porter).Error
	if err != nil {
		return nil, err
	}
	return &porter, nil
}

// FindByCode retrieves a porter from the database based on the given code.
// It returns a pointer to the porter and an error if any occurred.
func (r *PostgresOrmRepository) FindByCode(code domain.PorterCode) (*domain.Porter, error) {
	var porter domain.Porter
	err := r.db.Where("code = ?", code).First(&porter).Error
	if err != nil {
		return nil, err
	}
	return &porter, nil
}
