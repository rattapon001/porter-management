package postgresorm

import (
	infraErrors "github.com/rattapon001/porter-management/internal/infra/errors"
	"github.com/rattapon001/porter-management/internal/stock/domain"

	"gorm.io/gorm"
)

type ItemPostgresOrmRepository struct {
	db *gorm.DB
}

// NewPostgresOrmRepository creates a new instance of the PostgresOrmRepository.
// It takes a *gorm.DB as a parameter and returns a pointer to the PostgresOrmRepository.
func NewPostgresOrmRepository(db *gorm.DB) *ItemPostgresOrmRepository {
	return &ItemPostgresOrmRepository{
		db: db,
	}
}

func (r *ItemPostgresOrmRepository) Save(item *domain.Item) error {
	var existingItem domain.Item
	currentVersion := item.Version
	err := r.db.Where("id = ?", item.ID).First(&existingItem).Error
	if err == gorm.ErrRecordNotFound {
		// Job doesn't exist, insert it
		err = r.db.Create(item).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {

		if existingItem.Version != currentVersion {
			return infraErrors.ErrVersionMismatch
		}
		item.Version++
		err = r.db.Save(item).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ItemPostgresOrmRepository) FindById(id domain.ItemId) (*domain.Item, error) {
	var item domain.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemPostgresOrmRepository) FindAll() ([]*domain.Item, error) {
	var items []*domain.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
