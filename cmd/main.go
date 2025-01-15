package main

import (
	"invoice-system/config"
	"invoice-system/internal/database"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to PostgreSQL
	database.ConnectPostgres()

	// Connect to MongoDB
	database.ConnectMongo()

	// Initialize Echo instance
	e := echo.New()

	// Health check route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "up",
		})
	})

	// Start the server
	port := config.GetEnv("PORT", ":8080")
	log.Printf("Starting server on port %s\n", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
