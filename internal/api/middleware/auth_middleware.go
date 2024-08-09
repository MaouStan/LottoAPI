package middleware

import (
	"lottery-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		user, err := services.ValidateJWT(token)
		if err != nil || user.Role != role {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
