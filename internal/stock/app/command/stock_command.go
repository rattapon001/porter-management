package command

type StockCommand interface {
	Execute(eventName string, payload []byte)
}
