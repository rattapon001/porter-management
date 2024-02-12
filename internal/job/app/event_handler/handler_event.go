package event

import (
	"github.com/rattapon001/porter-management/pkg"
)

type EventHandler interface {
	Publish(events []pkg.Event) error
}
