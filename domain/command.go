package domain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wyllisMonteiro/event-sourcing/cqrs"
	"github.com/wyllisMonteiro/event-sourcing/models"
)

type CreateOrderCommand struct {
	models.Order
	Type string
}

type CreateOrderCommandHandler struct{}

func (ch CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {

	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		fmt.Println(cmd)
		orderCmd := command.Payload().(*CreateOrderCommand)
		order := &models.Order{
			People: orderCmd.People,
		}
		orderStr, _ := json.Marshal(order)
		event := &models.Event{
			Agg:     "order1",
			Handle:  orderCmd.Type,
			Payload: string(orderStr),
		}

		models.CreateOrder()

		models.CreateEvent(event)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
