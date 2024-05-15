package main

import (
	"database/sql"

	"github.com/MSPR-PayeTonKawa/orders/handlers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	db := sql.DB{}

	h := handlers.NewHandler(&db)

	// Define a route handler for the root path
	router.GET("/", h.HelloWorld)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start the server
	router.Run(":8080")
}
