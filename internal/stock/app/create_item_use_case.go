package app

import "github.com/rattapon001/porter-management/internal/stock/domain"

func (s *StockUseCaseImpl) CreateItem(item *domain.Item) (*domain.Item, error) {
	item, err := domain.NewItem(item.Name, item.Qty, item.Code)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(item); err != nil {
		return nil, err
	}
	return item, nil
}
