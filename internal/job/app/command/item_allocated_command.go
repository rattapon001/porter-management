package command

import (
	"context"

	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/rattapon001/porter-management/pkg"
)

type ItemAllocateCommand struct {
	JobsUseCase app.JobUseCase
}

type ItemAllocateEventPayload struct {
	ref   string
	items []domain.Equipment
}

func (i *ItemAllocateCommand) Execute(event interface{}) {
	if eventData, ok := event.(pkg.Event); ok {
		if eventPayload, ok := eventData.Payload.(ItemAllocateEventPayload); ok {
			i.JobsUseCase.JobAllocate(context.Background(), domain.JobId(eventPayload.ref), eventPayload.items)
		}
	}
}
