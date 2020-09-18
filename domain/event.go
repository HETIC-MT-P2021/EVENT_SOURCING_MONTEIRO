package domain

import (
	"reflect"

	"github.com/wyllisMonteiro/event-sourcing/models"
)

type HandleEvent interface {
	Execute()
}

type OrderCreated struct{}

func (event OrderCreated) Execute() {
}

var Events map[string]HandleEvent = map[string]HandleEvent{}

func InitEvents() {
	createOrderCmd := reflect.TypeOf(OrderCreated{}).String()
	Events[createOrderCmd] = OrderCreated{}
}

func LaunchEvents(events []models.Event) {
	for _, event := range events {
		if e, ok := Events[event.Handle]; ok {
			e.Execute()
		}
	}
}
