package command

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rattapon001/porter-management/internal/stock/app"
	"github.com/rattapon001/porter-management/internal/stock/domain"
)

type StockAllocateCommand struct {
	StockUseCase app.StockUseCase
}

type jobCreateEventPayloadEquipments struct {
	EquipmentId string `json:"EquipmentId"`
	Amount      int    `json:"Amount"`
	JobId       string `json:"JobId"`
}

type jobCreateEventPayload struct {
	Equipments []jobCreateEventPayloadEquipments `json:"equipments"`
	JobId      string                            `json:"jobId"`
}

func (s *StockAllocateCommand) Execute(eventName string, payload []byte) error {
	var data jobCreateEventPayload
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		fmt.Println("Error unmarshal payload : ", err)
		return err
	}
	items := []domain.Item{}
	for _, item := range data.Equipments {
		items = append(items, domain.Item{
			ID:  domain.ItemId(item.EquipmentId),
			Qty: item.Amount,
		})
	}
	_, err = s.StockUseCase.ItemAllocate(context.Background(), items, data.JobId)
	if err != nil {
		fmt.Println("Error allocate stock : ", err)
		return err
	}
	return nil
	// return errors.New("not implemented")
}
