package handlers

import (
	"log"
	"net/http"

	"github.com/MSPR-PayeTonKawa/orders/types"
	"github.com/gin-gonic/gin"
)

func (h Handlers) GetOrders(c *gin.Context) {
	rows, err := h.db.Query("SELECT * FROM orders")

	if err != nil {
		log.Fatal(err)
	}

	var orders []types.Order

	for rows.Next() {
		var OrderId string
		var CustomerName string
		var OrderDate string
		var Status string
		var Total float64
		if err := rows.Scan(&OrderId, &CustomerName, &OrderDate, &Status, &Total); err != nil {
			log.Fatal(err)
		}

		newOrder := types.Order{OrderId: OrderId, CustomerName: CustomerName, OrderDate: OrderDate, Status: Status, Total: Total}
		orders = append(orders, newOrder)
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func (h Handlers) AddOrder(c *gin.Context) {
	ct := c.Request.Header.Get("Content-Type")

	if ct != "application/json" {
		c.JSON(http.StatusInternalServerError, "Content-Type header is not application/json")
		return
	}

	var order types.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrderId := ""
	sqlStatement := "INSERT INTO orders (orderId, customerName, orderDate, status, total) VALUES ($1, $2, $3, $4, $5) RETURNING orderId"
	err := h.db.QueryRow(sqlStatement, order.OrderId, order.CustomerName, order.OrderDate, order.Status, order.Total).Scan(&newOrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orderId": newOrderId})
}

func (h Handlers) GetOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	row := h.db.QueryRow("SELECT orderId, customerName, orderDate, status, total FROM orders WHERE orderId = $1", orderId)

	var order types.Order
	err := row.Scan(&order.OrderId, &order.CustomerName, &order.OrderDate, &order.Status, &order.Total)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (h Handlers) ReplaceOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	ct := c.Request.Header.Get("Content-Type")

	if ct != "application/json" {
		c.JSON(http.StatusInternalServerError, "Content-Type header is not application/json")
		return
	}

	var updatedOrder types.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `
		UPDATE orders
		SET customerName = $1, orderDate = $2, status = $3, total = $4
		WHERE orderId = $5
	`
	_, err := h.db.Exec(sqlStatement, updatedOrder.CustomerName, updatedOrder.OrderDate, updatedOrder.Status, updatedOrder.Total, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "order updated"})
}

func (h Handlers) DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	sqlQuery := "DELETE FROM orders WHERE orderId = $1"
	_, err := h.db.Exec(sqlQuery, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted orderId": orderId})
}
