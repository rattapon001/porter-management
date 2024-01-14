package command

import (
	"fmt"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/pkg"
)

type PorterAllowcateCommand struct {
	PorterService app.PorterService
}

func (p *PorterAllowcateCommand) Execute(event interface{}) {
	if eventData, ok := event.(pkg.Event); ok {
		// Do something with eventPayload
		fmt.Printf("PorterAllowcateCommand: %v\n", eventData)
		if eventPayload, ok := eventData.Payload.(app.JobCreatedEvent); ok {
			fmt.Printf("PorterAllowcateCommand: %v\n", eventPayload)
		}
	}
}
