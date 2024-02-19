package command

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type ItemAllocateCommand struct {
	JobsUseCase app.JobUseCase
}

type AllocateEquipments struct {
	EquipmentId string `json:"EquipmentId"`
	Qty         int    `json:"Qty"`
}

type ItemAllocateEvent struct {
	JobId string               `json:"jobId"`
	Items []AllocateEquipments `json:"items"`
}

func (i *ItemAllocateCommand) Execute(eventName string, payload []byte) error {
	var data ItemAllocateEvent

	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		fmt.Println("Error unmarshal payload : ", err)
		return err
	}
	_, err = i.JobsUseCase.JobAllocate(context.Background(), domain.JobId(data.JobId))
	if err != nil {
		return err
	}

	return nil
}
