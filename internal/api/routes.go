package api

import (
	"lottery-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/", handlers.Hello)
}
