package domain

type PorterRepository interface {
	Save(porter *Porter) error
	FindAvailablePorter() (*Porter, error)
	FindByID(id PorterId) (*Porter, error)
	FindByCode(code PorterCode) (*Porter, error)
}
