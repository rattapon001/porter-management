package command

type PorterCommand interface {
	Execute(eventName string, payload []byte) error
}
