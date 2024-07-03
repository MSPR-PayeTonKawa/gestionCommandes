package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MSPR-PayeTonKawa/orders/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func simpleFunctionToTest(params string) string {
	return "Hello " + params
}

func connectDatabase() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	log.Println(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), fmt.Sprintf("%d", status)).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration.Seconds())
	}
}

func main() {
	if os.Getenv("GIN_MODE") != "release" {
		log.Println("Loading .env file")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Create a new Gin router
	router := gin.Default()

	// Apply the metrics middleware
	router.Use(metricsMiddleware())

	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(db)

	// Define a route handler for the metrics path
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	orders := router.Group("/orders")
	{
		orders.GET("/", h.GetOrders)
		orders.POST("/", h.AddOrder)
		orders.GET("/:orderId", h.GetOrder)
		orders.PUT("/:orderId", h.ReplaceOrder)
		orders.DELETE("/:orderId", h.DeleteOrder)
	}

	orderItem := router.Group("/item")
	{
		orderItem.GET("/", h.GetOrderItems)
		orderItem.POST("/", h.AddOrderItem)
		orderItem.GET("/:itemId", h.GetOrderItem)
		orderItem.PUT("/:itemId", h.ReplaceOrderItem)
		orderItem.DELETE("/:itemId", h.DeleteOrderItem)
		orderItem.DELETE("/clean", h.CleanOrderItem)
	}

	// Start the server
	router.Run(":8080")
}
