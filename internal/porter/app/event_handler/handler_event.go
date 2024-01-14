package event

import "github.com/rattapon001/porter-management/pkg"

type EventName string

type EventHandler interface {
	Publish(event []pkg.Event) error
}
