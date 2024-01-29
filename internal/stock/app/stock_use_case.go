package app

import (
	event "github.com/rattapon001/porter-management/internal/stock/app/event_handler"
	"github.com/rattapon001/porter-management/internal/stock/domain"
)

type StockUseCase interface {
	CreateItem(item *domain.Item) (*domain.Item, error)
	GetItem(id int) (*domain.Item, error)
	GetItems() ([]*domain.Item, error)
	ItemAllocate(items []domain.Item, consumerRef string) (*domain.Item, error)
}

type StockUseCaseImpl struct {
	repo      domain.ItemRepository
	Publisher event.EventHandler
}
