package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wyllisMonteiro/event-sourcing/models"
)

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	var product models.Product
	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		fmt.Println(err)
	}

	product_id := models.CreateProduct(product)

	fmt.Println(product_id)
}
