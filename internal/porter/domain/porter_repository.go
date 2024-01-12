package domain

type PorterRepository interface {
	Save(porter *Porter) error
}
