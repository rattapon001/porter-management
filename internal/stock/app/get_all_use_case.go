package app

import "github.com/rattapon001/porter-management/internal/stock/domain"

func (s *StockUseCaseImpl) GetItems() ([]*domain.Item, error) {
	items, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}
