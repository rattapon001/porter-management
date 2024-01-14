package app

type JobCreatedEvent struct {
	JobId string
}

func (s *PorterServiceImpl) PorterAllowcate(payload JobCreatedEvent) error {
	availablePorter := s.Repo.FindAvailablePorter()
	if availablePorter == nil {
		return nil
	}
	s.Noti.Notify(availablePorter.Token, payload.JobId)

	return nil
}
