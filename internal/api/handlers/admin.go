package handlers

import (
	"lottery-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DrawLottoNumbers handles the drawing of lotto numbers by the admin
func DrawLottoNumbers(c *gin.Context) {
	count := 5 // You can adjust this as needed or retrieve it from the request body
	numbers, err := services.GenerateRandomNumbers(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to draw lotto numbers"})
		return
	}

	// Store the drawn numbers in the database
	err = services.StoreDrawnNumbers(numbers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store drawn numbers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"winning_numbers": numbers})
}

// ResetSystem handles the system reset by the admin
func ResetSystem(c *gin.Context) {
	err := services.ResetSystem()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "System reset successfully"})
}
