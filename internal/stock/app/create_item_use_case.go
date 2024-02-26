package app

import (
	"context"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/internal/stock/domain"
)

func (s *StockUseCaseImpl) CreateItem(ctx context.Context, item *domain.Item) (*domain.Item, error) {
	var result *domain.Item
	s.Uow.DoInTx(ctx, func(store uow.UnitOfWorkStore) error {
		item, err := domain.NewItem(item.Name, item.Qty, item.Sku)
		if err != nil {
			return err
		}
		if err := s.repo.Save(item); err != nil {
			return err
		}
		if err := s.Publisher.Publish(item.Aggregate.Events, store); err != nil {
			return err
		}
		result = item
		return nil
	})

	return result, nil
}
