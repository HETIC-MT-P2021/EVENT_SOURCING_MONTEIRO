package models

type Order struct {
	ID     int    `json:"id"`
	People string `json:"people"`
}

type OrderReq struct {
	ID       int    `json:"id"`
	People   string `json:"people"`
	Products []int  `json:"products"`
}

func CreateOrder(orderReq interface{}) int {
	var order = Order{
		People: orderReq.(OrderReq).People,
	}

	DB.AutoMigrate(&Order{})

	DB.Save(&order)

	return order.ID
}
