package app

import "github.com/rattapon001/porter-management/pkg"

type NotificationUseCase interface {
	Notify(token string, payload pkg.NotificationPayload) error
}

type NotificationUseCaseImpl struct {
}

func NewNotificationUseCase() NotificationUseCase {
	return &NotificationUseCaseImpl{}
}
