package app

import (
	"context"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	event "github.com/rattapon001/porter-management/internal/stock/app/event_handler"
	"github.com/rattapon001/porter-management/internal/stock/domain"
)

type StockUseCase interface {
	CreateItem(item *domain.Item) (*domain.Item, error)
	GetItem(id domain.ItemId) (*domain.Item, error)
	GetItems() ([]*domain.Item, error)
	ItemAllocate(ctx context.Context, items []domain.Item, consumerRef string) (*domain.Item, error)
}

type StockUseCaseImpl struct {
	repo      domain.ItemRepository
	Publisher event.EventHandler
	Uow       uow.UnitOfWork
}

func NewStockUseCase(repo domain.ItemRepository, publisher event.EventHandler, uow uow.UnitOfWork) StockUseCase {
	return &StockUseCaseImpl{
		repo:      repo,
		Publisher: publisher,
		Uow:       uow,
	}
}
