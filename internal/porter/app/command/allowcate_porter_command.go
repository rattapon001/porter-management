package command

import (
	"encoding/json"
	"fmt"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterAllocateCommand struct {
	PorterUseCase app.PorterUseCase
}

type jobAllocatedEventPayload struct {
	JobId    string          `json:"jobId"`
	Status   string          `json:"status"`
	Location domain.Location `json:"locations"`
	Patient  domain.Patient  `json:"patient"`
}

func (p *PorterAllocateCommand) Execute(eventName string, payload []byte) error {

	var data jobAllocatedEventPayload
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		fmt.Println("Error unmarshal payload : ", err)
		return err
	}

	job := domain.NewJob(domain.JobId(data.JobId), data.Patient, data.Location)

	_, err = p.PorterUseCase.PorterAllocate(*job)
	if err != nil {
		fmt.Println("Error allocate porter) : ", err)
		return err
	}
	return nil
}
