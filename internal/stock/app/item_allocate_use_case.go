package app

import (
	"github.com/rattapon001/porter-management/internal/stock/domain"
	"github.com/rattapon001/porter-management/pkg"
)

const (
	ItemAllocatedEventName pkg.EventName = "ItemAllocated"
)

func (s *StockUseCaseImpl) ItemAllocate(items []domain.Item, consumerRef string) (*domain.Item, error) {

	ItemEventPayload := []map[string]interface{}{}
	for _, item := range items {
		itemOne, err := s.repo.FindById(item.ID)
		if err != nil {
			return nil, err
		}
		err = itemOne.ItemAllocate(item.Qty)
		if err != nil {
			return nil, err
		}
		if err := s.repo.Save(itemOne); err != nil {
			return nil, err
		}
		ItemEventPayload = append(ItemEventPayload, map[string]interface{}{
			"EquipmentId": itemOne.ID,
			"Qty":         itemOne.Qty,
		})
	}
	if err := s.Publisher.Publish([]pkg.Event{
		{
			EventName: ItemAllocatedEventName,
			Payload: map[string]interface{}{
				"ref":   consumerRef,
				"items": ItemEventPayload,
			},
		},
	}); err != nil {
		return nil, err
	}

	return nil, nil

}
