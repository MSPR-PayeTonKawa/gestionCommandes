package main

import (
	"database/sql"

	"github.com/MSPR-PayeTonKawa/orders/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	db := sql.DB{}

	h := handlers.NewHandler(&db)

	// Define a route handler for the root path
	router.GET("/", h.HelloWorld)

	// Start the server
	router.Run(":8080")
}
