package domain

type PorterRepository interface {
	Save(porter *Porter) error
	Update(porter *Porter) error
	FindAvailablePorter() *Porter
	FindByID(id PorterId) *Porter
	FindByCode(code PorterCode) *Porter
}
