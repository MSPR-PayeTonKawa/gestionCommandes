package types

type Order struct {
	OrderId      string  `json:"orderId" example:"12345"`
	CustomerName string  `json:"customerName" example:"John Doe"`
	OrderDate    string  `json:"orderDate" format:"date-time" example:"2021-06-23T18:25:43.511Z"`
	Status       string  `json:"status" example:"Shipped"`
	Total        float64 `json:"total" example:"150.50"`
}
