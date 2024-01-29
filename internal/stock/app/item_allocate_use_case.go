package app

import "github.com/rattapon001/porter-management/internal/stock/domain"

func (s *StockUseCaseImpl) ItemAllocate(id string, qty int, consumerRef string) (*domain.Item, error) {
	item, err := s.repo.FindById(domain.ItemId(id))
	if err != nil {
		return nil, err
	}
	err = item.ItemAllocate(qty, domain.Consumer{Ref: consumerRef})
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(item); err != nil {
		return nil, err
	}
	if err := s.Publisher.Publish(item.Aggregate.Events); err != nil {
		return nil, err
	}
	return item, nil
}
