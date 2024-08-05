package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maoustan/lotto-api/internal/api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/lotto")
	// auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", handlers.Hello)
		// เพิ่ม routes อื่นๆ ตามต้องการ
	}

	return r
}
