package types

type OrderItem struct {
	ProductId string  `json:"productId" example:"P100"`
	Quantity  int     `json:"quantity" example:"2"`
	Price     float64 `json:"price" format:"float" example:"25.50"`
}
