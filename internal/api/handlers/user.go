package handlers

import (
	"lottery-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWalletBalance(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	balance, err := services.GetWalletBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get wallet balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
