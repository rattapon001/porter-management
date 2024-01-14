package command

type PorterCommand interface {
	Execute(event ...interface{}) error
}
