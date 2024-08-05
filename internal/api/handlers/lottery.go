package handlers

import (
	"lottery-api/internal/models"
	"lottery-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLotteryNumbers(c *gin.Context) {
	numbers, err := services.GetLotteryNumbers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get lottery numbers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numbers": numbers})
}

func PurchaseLottery(c *gin.Context) {
	var req models.PurchaseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := services.PurchaseLottery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to purchase lottery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lottery purchased successfully"})
}

func TransferLottery(c *gin.Context) {
	var req models.TransferRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := services.TransferLottery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transfer lottery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lottery transferred successfully"})
}
