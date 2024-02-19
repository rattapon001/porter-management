package command

type JobCommand interface {
	Execute(eventName string, payload []byte) error
}
