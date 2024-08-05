package api

import (
	"lottery-api/internal/api/handlers"
	// "lottery-api/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/", handlers.Hello)
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Middleware for authentication
	authRoutes := router.Group("/")
	// authRoutes.Use(middleware.Auth())

	// Authenticated routes
	authRoutes.GET("/lottery", handlers.GetLotteryNumbers)
	authRoutes.POST("/purchase", handlers.PurchaseLottery)
	authRoutes.POST("/transfer", handlers.TransferLottery)
	authRoutes.GET("/wallet", handlers.GetWalletBalance)
	authRoutes.GET("/draw", handlers.GetDrawResults)
}
