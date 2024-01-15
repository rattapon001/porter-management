package app

import "github.com/rattapon001/porter-management/pkg"

type NotificationService interface {
	Notify(token string, payload pkg.NotificationPayload) error
}

type NotificationServiceImpl struct {
}

func NewNotificationService() NotificationService {
	return &NotificationServiceImpl{}
}
