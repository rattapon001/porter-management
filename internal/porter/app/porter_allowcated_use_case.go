package app

import (
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/rattapon001/porter-management/pkg"
)

func (s *PorterServiceImpl) PorterAllowcated(payload domain.Job) (*domain.Porter, error) {
	availablePorter := s.Repo.FindAvailablePorter()
	if availablePorter == nil {
		return nil, nil
	}
	NotiPayload := pkg.NotificationPayload{
		JobId:   string(payload.ID),
		Version: payload.Version,
		Message: "Job created " + payload.Location.From + " to " + payload.Location.To + " for " + payload.Patient.Name,
	}
	if err := s.Noti.Notify(availablePorter.Token, NotiPayload); err != nil {
		return nil, err
	}
	return availablePorter, nil
}
