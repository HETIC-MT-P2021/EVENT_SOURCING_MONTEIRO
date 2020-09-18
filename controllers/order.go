package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/wyllisMonteiro/event-sourcing/cqrs"
	"github.com/wyllisMonteiro/event-sourcing/domain"
	"github.com/wyllisMonteiro/event-sourcing/models"
)

func CreateOrder(w http.ResponseWriter, req *http.Request) {
	var orderReq models.OrderReq
	if err := json.NewDecoder(req.Body).Decode(&orderReq); err != nil {
		fmt.Println(err)
		return
	}

	domain.InitBusses()

	command := cqrs.NewCommandMessage(&domain.CreateOrderCommand{
		models.Order{
			People: orderReq.People,
		},
		reflect.TypeOf(domain.OrderCreated{}).String(),
	})
	_ = domain.CommandBus.Dispatch(command)

}

func LaunchEvents(w http.ResponseWriter, req *http.Request) {
	domain.InitEvents()
	events := models.GetEvents()
	domain.LaunchEvents(events)
}
