package handlers

import (
	"log"

	"github.com/MSPR-PayeTonKawa/orders/types"
	"github.com/gin-gonic/gin"
)

func (h Handlers) GetOrders(c *gin.Context) {
	rows, err := h.db.Query("SELECT * FROM orders")

	var orders []types.OrderItem

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ProductId string
		var Quantity int
		var Price float64
		if err := rows.Scan(&ProductId, &Quantity, &Price); err != nil {
			log.Fatal(err)
		}

		newOrder := types.OrderItem{ProductId: ProductId, Quantity: Quantity, Price: Price}
		orders = append(orders, newOrder)
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})

}

func (h Handlers) AddOrder(c *gin.Context)     {}
func (h Handlers) GetOrder(c *gin.Context)     {}
func (h Handlers) ReplaceOrder(c *gin.Context) {}
func (h Handlers) DeleteOrder(c *gin.Context)  {}
