package services

import (
	"errors"
	"lottery-api/internal/db"
	"lottery-api/internal/models"
)

// GetLotteryNumbers retrieves all lottery numbers from the database.
func GetLotteryNumbers() ([]models.LotteryNumber, error) {
	rows, err := db.DB.Query("SELECT number_id, number_value FROM lottery_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []models.LotteryNumber
	for rows.Next() {
		var number models.LotteryNumber
		if err := rows.Scan(&number.NumberID, &number.NumberValue); err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// PurchaseLottery processes a lottery ticket purchase.
func PurchaseLottery(req models.PurchaseRequest) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	// Check member balance
	var balance float64
	err = tx.QueryRow("SELECT wallet_balance FROM members WHERE member_id = $1", req.MemberID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return err
	}

	if balance < req.Amount {
		tx.Rollback()
		return errors.New("insufficient funds")
	}

	// Deduct balance
	_, err = tx.Exec("UPDATE members SET wallet_balance = wallet_balance - $1 WHERE member_id = $2", req.Amount, req.MemberID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Record purchase
	_, err = tx.Exec("INSERT INTO purchases (member_id, number_id, purchase_date) VALUES ($1, $2, CURRENT_TIMESTAMP)", req.MemberID, req.NumberID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// TransferLottery processes the transfer of a lottery ticket from one member to another.
func TransferLottery(req models.TransferRequest) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	// Check if the sender has the lottery ticket
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM purchases WHERE member_id = $1 AND number_id = $2", req.FromMemberID, req.NumberID).Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("ticket not found")
	}

	// Record transfer
	_, err = tx.Exec("INSERT INTO transfers (from_member_id, to_member_id, number_id, transfer_date) VALUES ($1, $2, $3, CURRENT_TIMESTAMP)", req.FromMemberID, req.ToMemberID, req.NumberID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Remove the ticket from the sender's purchases
	_, err = tx.Exec("DELETE FROM purchases WHERE member_id = $1 AND number_id = $2", req.FromMemberID, req.NumberID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add the ticket to the receiver's purchases
	_, err = tx.Exec("INSERT INTO purchases (member_id, number_id, purchase_date) VALUES ($1, $2, CURRENT_TIMESTAMP)", req.ToMemberID, req.NumberID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
