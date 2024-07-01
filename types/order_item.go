package types

type OrderItem struct {
	ItemId    int     `json:"itemId" example:"12345"`
	OrderId   string  `json:"orderId" example:"12345"`
	ProductId string  `json:"productId" example:"P100"`
	Quantity  int     `json:"quantity" example:"2"`
	Price     float64 `json:"price" format:"float" example:"25.50"`
}
