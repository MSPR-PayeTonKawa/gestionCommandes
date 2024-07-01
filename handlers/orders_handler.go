package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MSPR-PayeTonKawa/orders/types"
	"github.com/gin-gonic/gin"
)

func (h Handlers) GetOrders(c *gin.Context) {
	//log.Print("GetOrders")
	rows, err := h.db.Query("SELECT * FROM orders")

	var orders []types.Order

	if err != nil {
		log.Fatal(err)
	}

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

	dec := json.NewDecoder(c.Request.Body)
	log.Print("AddOrder : ", &dec)

	var order *types.Order
	err := dec.Decode(&order)

	if err != nil {
		log.Fatal(err)
	}

	newOrderId := ""
	sqlStatement := "INSERT INTO orders (orderId, customerName, orderDate, status, total) VALUES ($1, $2, $3, $4, $5) RETURNING orderId"
	err = h.db.QueryRow(sqlStatement, order.OrderId, order.CustomerName, order.OrderDate, order.Status, order.Total).Scan(&newOrderId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"itemId": newOrderId})
}
func (h Handlers) GetOrder(c *gin.Context)     {}
func (h Handlers) ReplaceOrder(c *gin.Context) {}
func (h Handlers) DeleteOrder(c *gin.Context)  {}
