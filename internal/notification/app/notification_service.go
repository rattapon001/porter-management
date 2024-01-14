package app

type NotificationService interface {
	Notify(token, payload string) error
}

type NotificationServiceImpl struct {
}

func NewNotificationService() NotificationService {
	return &NotificationServiceImpl{}
}
