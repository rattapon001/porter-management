package app

import "github.com/rattapon001/porter-management/pkg"

type JobCreatedEvent struct {
	JobId string
}

func (s *PorterServiceImpl) PorterAllowcated(payload JobCreatedEvent) error {
	availablePorter := s.Repo.FindAvailablePorter()
	if availablePorter == nil {
		return nil
	}

	NotiPayload := pkg.NotificationPayload{
		JobId:   payload.JobId,
		Message: "Your job is ready",
	}

	s.Noti.Notify(availablePorter.Token, NotiPayload)
	return nil
}
