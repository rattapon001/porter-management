package app

type NotificationService interface {
	Notify(token, message string) error
}

type NotificationServiceImpl struct {
}

func NewNotificationService() NotificationService {
	return &NotificationServiceImpl{}
}
