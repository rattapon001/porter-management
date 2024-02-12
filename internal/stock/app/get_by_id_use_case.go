package app

import "github.com/rattapon001/porter-management/internal/stock/domain"

func (s *StockUseCaseImpl) GetItem(id domain.ItemId) (*domain.Item, error) {
	item, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
