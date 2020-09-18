package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/event-sourcing/models"
	"github.com/wyllisMonteiro/event-sourcing/router"
)

func main() {
	models.ConnectToBDD()

	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))

	models.DB.Close()
}
