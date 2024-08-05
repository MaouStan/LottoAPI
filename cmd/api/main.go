package main

import (
	"lottery-api/internal/api"
	"lottery-api/internal/api/middleware"
	"lottery-api/internal/config"
	"lottery-api/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.Load()

	// Initialize database
	db.Connect()

	// Initialize Gin router
	router := gin.Default()

	// Middleware
	router.Use(middleware.Logging())

	// Set up routes
	api.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
