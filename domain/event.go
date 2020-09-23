package domain

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/wyllisMonteiro/event-sourcing/cqrs"
	"github.com/wyllisMonteiro/event-sourcing/models"
)

type HandleEvent interface {
	Execute(string)
}

type OrderCreated struct{}

func (event OrderCreated) Execute(payload string) {
	bytes := []byte(payload)
	var order models.Order
	json.Unmarshal(bytes, &order)

	fmt.Println(order.People)

	command := cqrs.NewCommandMessage(&CreateOrderCommand{
		models.Order{
			People: order.People,
		},
		reflect.TypeOf(OrderCreated{}).String(),
	})
	_ = CommandBus.Dispatch(command)
}

var Events map[string]HandleEvent = map[string]HandleEvent{}

func InitEvents() {
	createOrderCmd := reflect.TypeOf(OrderCreated{}).String()
	Events[createOrderCmd] = OrderCreated{}
}

func LaunchEvents(events []models.Event) {
	for _, event := range events {
		if e, ok := Events[event.Handle]; ok {
			e.Execute(event.Payload)
		}
	}
}
