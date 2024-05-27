package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MSPR-PayeTonKawa/orders/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(db)

	// Define a route handler for the root path
	router.GET("/", h.HelloWorld)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start the server
	router.Run(":8080")
}
