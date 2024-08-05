package handlers

import (
	"lottery-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDrawResults(c *gin.Context) {
	results, err := services.GetDrawResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get draw results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
