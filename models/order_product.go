package models

type OrderProduct struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
}

type OrderProductReq struct {
	OrderID    int   `json:"order_id"`
	ProductIDs []int `json:"product_ids"`
}

func CreateOrderProduct(orderProductReq OrderProductReq) {

	DB.AutoMigrate(&OrderProduct{})

	for _, productID := range orderProductReq.ProductIDs {
		orderProduct := OrderProduct{
			OrderID:   orderProductReq.OrderID,
			ProductID: productID,
		}
		DB.Save(&orderProduct)
	}
}
