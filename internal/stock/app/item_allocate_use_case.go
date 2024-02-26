package app

import (
	"context"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/internal/stock/domain"
	"github.com/rattapon001/porter-management/pkg"
)

const (
	ItemAllocatedEventName pkg.EventName = "ItemAllocated"
)

func (s *StockUseCaseImpl) ItemAllocate(ctx context.Context, items []domain.Item, jobId string) (*domain.Item, error) {
	ItemEventPayload := []map[string]interface{}{}
	err := s.Uow.DoInTx(ctx, func(store uow.UnitOfWorkStore) error {
		for _, item := range items {
			itemOne, err := s.repo.FindById(item.ID)

			if err != nil {
				return err
			}
			err = itemOne.ItemAllocate(item.Qty)
			if err != nil {
				return err
			}
			if err := store.Item().Save(itemOne); err != nil {
				return err
			}
			ItemEventPayload = append(ItemEventPayload, map[string]interface{}{
				"EquipmentId": itemOne.ID,
				"Qty":         itemOne.Qty,
			})
		}

		if err := s.Publisher.Publish([]pkg.Event{
			{
				EventName: domain.ItemAllocatedEvent,
				Payload: map[string]interface{}{
					"jobId": jobId,
					"items": ItemEventPayload,
				},
			},
		}, store); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil

}
