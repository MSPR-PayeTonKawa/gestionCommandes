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

	var userOrder *types.OrderItem
	err := dec.Decode(&userOrder)

	if err != nil {
		log.Fatal(err)
	}

	newItemId := -1
	sqlStatement := "INSERT INTO OrderItems (productId, quantity, price) VALUES ($1, $2, $3) RETURNING itemId"
	err = h.db.QueryRow(sqlStatement, userOrder.ProductId, userOrder.Quantity, userOrder.Price).Scan(&newItemId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"itemId": newItemId})
}
func (h Handlers) GetOrder(c *gin.Context)     {}
func (h Handlers) ReplaceOrder(c *gin.Context) {}
func (h Handlers) DeleteOrder(c *gin.Context)  {}
