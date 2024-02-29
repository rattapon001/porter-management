package app

import (
	"log"

	"github.com/rattapon001/porter-management/pkg"
)

func (m *NotificationUseCaseImpl) Notify(token string, payload pkg.NotificationPayload) error {
	log.Printf("send notification to %s with payload %+v\n", token, payload)
	return nil
}
