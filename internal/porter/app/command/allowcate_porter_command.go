package command

import (
	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/pkg"
)

type PorterAllowcateCommand struct {
	PorterService app.PorterService
}

func (p *PorterAllowcateCommand) Execute(event interface{}) {
	if eventData, ok := event.(pkg.Event); ok {
		if eventPayload, ok := eventData.Payload.(app.JobCreatedEvent); ok {
			p.PorterService.PorterAllowcated(eventPayload)
		}
	}
}
