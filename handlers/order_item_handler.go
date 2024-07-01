package handlers

import (
	"log"
	"net/http"

	"github.com/MSPR-PayeTonKawa/orders/types"
	"github.com/gin-gonic/gin"
)

func (h Handlers) GetOrderItems(c *gin.Context) {
	rows, err := h.db.Query("SELECT * FROM orderItems")

	if err != nil {
		log.Fatal(err)
	}

	var orders []types.OrderItem

	for rows.Next() {
		var ItemId int
		var OrderId string
		var ProductId string
		var Quantity int
		var Price float64
		if err := rows.Scan(&ItemId, &OrderId, &ProductId, &Quantity, &Price); err != nil {
			log.Fatal(err)
		}

		newOrder := types.OrderItem{ItemId: ItemId, OrderId: OrderId, ProductId: ProductId, Quantity: Quantity, Price: Price}
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

	var userOrder types.OrderItem
	if err := c.ShouldBindJSON(&userOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Print(userOrder, userOrder.OrderId)

	newItemId := -1
	sqlStatement := "INSERT INTO orderItems (orderId, productId, quantity, price) VALUES ($1, $2, $3, $4) RETURNING itemId"
	err := h.db.QueryRow(sqlStatement, userOrder.OrderId, userOrder.ProductId, userOrder.Quantity, userOrder.Price).Scan(&newItemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"itemId": newItemId})
}

func (h Handlers) GetOrderItem(c *gin.Context) {
	itemId := c.Param("itemId")

	log.Print("itemId : ", itemId)
	row := h.db.QueryRow("SELECT itemId, orderId, productId, quantity, price FROM orderItems WHERE itemId = $1", itemId)

	var item types.OrderItem
	err := row.Scan(&item.ItemId, &item.OrderId, &item.ProductId, &item.Quantity, &item.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

func (h Handlers) ReplaceOrderItem(c *gin.Context) {
	itemId := c.Param("itemId")

	ct := c.Request.Header.Get("Content-Type")

	if ct != "application/json" {
		c.JSON(http.StatusInternalServerError, "Content-Type header is not application/json")
		return
	}

	var updatedOrderItem types.OrderItem
	if err := c.ShouldBindJSON(&updatedOrderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `
		UPDATE orderItems
		SET orderId = $1, productId = $2, quantity = $3, price = $4
		WHERE itemId = $5
	`
	_, err := h.db.Exec(sqlStatement, updatedOrderItem.OrderId, updatedOrderItem.ProductId, updatedOrderItem.Quantity, updatedOrderItem.Price, itemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "order item updated"})
}

func (h Handlers) DeleteOrderItem(c *gin.Context) {
	itemId := c.Param("itemId")

	sqlQuery := "DELETE FROM orderItems WHERE itemId = $1"
	_, err := h.db.Exec(sqlQuery, itemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted itemId": itemId})
}

// TODO: clean this endpoint later
func (h Handlers) CleanOrderItem(c *gin.Context) {
	h.db.Query("DELETE FROM orderItems")
	c.Status(http.StatusOK)
}
