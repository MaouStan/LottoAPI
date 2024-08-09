package api

import (
	"lottery-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Hello World
	router.GET("/", handlers.Hello)

	// API routes
	api := router.Group("/api")
	{
		// User Routes
		userGroup := api.Group("/auth")
		{
			userGroup.POST("/register", handlers.SignUp) // User registration
			userGroup.POST("/login", handlers.Login)     // User login
		}
		// Member Routes (requires authentication)
		memberGroup := api.Group("/")
		// memberGroup.Use(middleware.AuthMiddleware("member"))
		{
			memberGroup.GET("/lotto", handlers.ViewLottoNumbers)     // View  lotto numbers
			memberGroup.POST("/purchase", handlers.PurchaseLotto)    // Purchase lotto number
			memberGroup.GET("/balance", handlers.CheckWalletBalance) // Check wallet balance
			memberGroup.GET("/results", handlers.CheckLottoResults)  // Check lotto results
			memberGroup.POST("/claim", handlers.ClaimWinnings)       // Claim winnings
		}

		// Admin Routes (requires authentication)
		adminGroup := api.Group("/admin")
		// adminGroup.Use(middleware.AuthMiddleware("admin"))
		{
			adminGroup.POST("/draw", handlers.DrawLottoNumbers) // Draw lotto numbers
			adminGroup.POST("/reset", handlers.ResetSystem)     // Reset the system
		}
	}
}
