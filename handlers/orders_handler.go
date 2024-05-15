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

	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})

}
