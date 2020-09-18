package router

import (
	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/event-sourcing/controllers"
)

// Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/", controllers.CreateOrder).Methods("POST")
	r.HandleFunc("/events", controllers.LaunchEvents).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	return r
}
