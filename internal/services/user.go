package services

import (
	"database/sql"
	"lottery-api/internal/db"
)

func GetWalletBalance(userID int) (float64, error) {
	var balance float64
	err := db.DB.QueryRow("SELECT wallet_balance FROM members WHERE member_id = $1", userID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return balance, nil
}
