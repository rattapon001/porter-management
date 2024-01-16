package command

import (
	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/rattapon001/porter-management/pkg"
)

type PorterAvailableCommand struct {
	PorterUseCase app.PorterUseCase
}

func (p *PorterAvailableCommand) Execute(event interface{}) {
	if eventData, ok := event.(pkg.Event); ok {
		if eventPayload, ok := eventData.Payload.(domain.Job); ok {
			p.PorterUseCase.PorterAvailable(domain.PorterCode(eventPayload.Porter.Code))
		}
	}
}
