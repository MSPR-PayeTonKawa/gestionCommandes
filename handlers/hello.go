package handlers

import "github.com/gin-gonic/gin"

func (h Handlers) HelloWorld(c *gin.Context) {
	h.db.Query("SELECT * FROM orders")

	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
