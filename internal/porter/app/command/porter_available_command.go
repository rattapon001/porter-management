package command

import (
	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/rattapon001/porter-management/pkg"
)

type PorterAvailableCommand struct {
	PorterService app.PorterService
}

func (p *PorterAvailableCommand) Execute(event interface{}) {
	if eventData, ok := event.(pkg.Event); ok {
		if eventPayload, ok := eventData.Payload.(domain.Job); ok {
			p.PorterService.PorterAvailable(domain.PorterId(eventPayload.Porter.Code))
		}
	}
}
