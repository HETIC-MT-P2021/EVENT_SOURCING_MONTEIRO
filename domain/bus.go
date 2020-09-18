package domain

import (
	"github.com/wyllisMonteiro/event-sourcing/cqrs"
)

var CommandBus *cqrs.CommandBus

func InitBusses() {
	CommandBus = cqrs.InitCommandBus()
	_ = CommandBus.RegisterHandler(NewCreateOrderCommandHandler(), &CreateOrderCommand{})
}
