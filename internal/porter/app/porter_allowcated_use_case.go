package app

type JobCreatedEvent struct {
	JobId string
}

func (s *PorterServiceImpl) PorterAllowcated(payload JobCreatedEvent) error {
	availablePorter := s.Repo.FindAvailablePorter()
	if availablePorter == nil {
		return nil
	}
	s.Noti.Notify(availablePorter.Token, payload.JobId)
	return nil
}
