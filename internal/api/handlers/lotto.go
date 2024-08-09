package handlers

import (
	"lottery-api/internal/models"
	"lottery-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ViewLottoNumbers handles viewing  lotto numbers
func ViewLottoNumbers(c *gin.Context) {
	numbers, err := services.GetAllLottoNumbers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lotto numbers"})
		return
	}
	c.JSON(http.StatusOK, numbers)
}

// PurchaseLotto handles purchasing a lotto number
func PurchaseLotto(c *gin.Context) {
	user, _ := c.Get("user") // Get the user from context
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var purchaseData struct {
		LottoNumberID int `json:"lotto_number_id"`
		Amount        int `json:"amount"`
	}
	if err := c.ShouldBindJSON(&purchaseData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := services.PurchaseLottoNumber(user.(*models.User).ID, purchaseData.LottoNumberID, purchaseData.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to purchase lotto number"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lotto number purchased successfully"})
}

// CheckWalletBalance handles checking wallet balance
func CheckWalletBalance(c *gin.Context) {
	user, _ := c.Get("user") // Get the user from context
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"wallet_balance": user.(*models.User).WalletBalance})
}

// CheckLottoResults handles checking lotto results
func CheckLottoResults(c *gin.Context) {
	results, err := services.GetDraw()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lotto results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ClaimWinnings handles claiming winnings
func ClaimWinnings(c *gin.Context) {
	user, _ := c.Get("user") // Get the user from context
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := services.ClaimWinnings(user.(*models.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to claim winnings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Winnings claimed successfully"})
}
