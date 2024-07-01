package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MSPR-PayeTonKawa/orders/types"
	"github.com/gin-gonic/gin"
)

func (h Handlers) GetOrderItems(c *gin.Context) {
	rows, err := h.db.Query("SELECT * FROM orderItems")

	log.Print(rows)

	var orders []types.OrderItem

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ItemId string
		var OrderId string
		var ProductId string
		var Quantity int
		var Price float64
		if err := rows.Scan(&ItemId, &OrderId, &ProductId, &Quantity, &Price); err != nil {
			log.Fatal(err)
		}

		newOrder := types.OrderItem{ProductId: ProductId, Quantity: Quantity, Price: Price}
		orders = append(orders, newOrder)
	}

	c.JSON(http.StatusOK, gin.H{
		"orderItems": orders,
	})
}

func (h Handlers) AddOrderItem(c *gin.Context) {
	ct := c.Request.Header.Get("Content-Type")

	if ct != "application/json" {
		c.JSON(http.StatusInternalServerError, "Content-Type header is not application/json")
		return
	}

	dec := json.NewDecoder(c.Request.Body)
	log.Print("AddOrderItem : ", &dec)

	var userOrder *types.OrderItem
	err := dec.Decode(&userOrder)

	if err != nil {
		log.Fatal(err)
	}

	newItemId := -1
	sqlStatement := "INSERT INTO orderItems (itemId, orderId, productId, quantity, price) VALUES ($1, $2, $3, $4, $5) RETURNING itemId"
	err = h.db.QueryRow(sqlStatement, userOrder.ItemId, userOrder.OrderId, userOrder.ProductId, userOrder.Quantity, userOrder.Price).Scan(&newItemId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"itemId": newItemId})
}

func (h Handlers) GetOrderItem(c *gin.Context)     {}
func (h Handlers) ReplaceOrderItem(c *gin.Context) {}
func (h Handlers) DeleteOrderItem(c *gin.Context)  {}

// TODO: clean this endpoint later
func (h Handlers) CleanOrderItem(c *gin.Context) {
	h.db.Query("DELETE FROM orderItems")
	c.Status(http.StatusOK)
}
