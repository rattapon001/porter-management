package command

import (
	"fmt"
)

type StockAllocateCommand struct {
	// StockUseCase app.StockUseCase
}

func (s *StockAllocateCommand) Execute(eventName string, payload interface{}) {
	fmt.Printf("StockAllocateCommand : %s\n", eventName)
	// if eventData, ok := event.(app.StockEvent); ok {
	// 	s.StockUseCase.StockAllocate(eventData)
	// }
}
