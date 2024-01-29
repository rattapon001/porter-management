package memory

import (
	infraErrors "github.com/rattapon001/porter-management/internal/infra/errors"
	"github.com/rattapon001/porter-management/internal/stock/domain"
)

type ItemMemoryRepository struct {
	items []*domain.Item
}

func NewItemMemoryRepository() *ItemMemoryRepository {
	return &ItemMemoryRepository{
		items: []*domain.Item{},
	}
}

func (r *ItemMemoryRepository) Save(item *domain.Item) error {
	for i, existingItem := range r.items {
		if existingItem.ID == item.ID {
			if existingItem.Version != item.Version {
				return infraErrors.ErrVersionMismatch
			}
			item.Version++
			r.items[i] = item
			return nil
		}
	}
	r.items = append(r.items, item)
	return nil
}

func (r *ItemMemoryRepository) FindById(id domain.ItemId) (*domain.Item, error) {
	for _, item := range r.items {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, nil
}
