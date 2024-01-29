package app

import "github.com/rattapon001/porter-management/internal/stock/domain"

type StockUseCase interface {
	CreateItem(item *domain.Item) (*domain.Item, error)
	GetItem(id int) (*domain.Item, error)
	GetItems() ([]*domain.Item, error)
	ItemAllocate(id int, quantity int, consumerRef string) (*domain.Item, error)
}

type StockUseCaseImpl struct {
	repo domain.ItemRepository
}
