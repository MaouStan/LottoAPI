package services

import (
	"errors"
	"lottery-api/internal/db"
	"lottery-api/internal/models"
	"math/rand"
	"time"
)

// GetAllLottoNumbers retrieves all  lotto numbers from the database
func GetAllLottoNumbers() ([]models.LottoNumber, error) {
	rows, err := db.Conn.Query(`
		SELECT id, number FROM lotto_numbers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []models.LottoNumber
	for rows.Next() {
		var num models.LottoNumber
		if err := rows.Scan(&num.ID, &num.Number); err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// PurchaseLottoNumber processes a lotto number purchase
func PurchaseLottoNumber(userID, lottoNumberID int, amount int) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var walletBalance int
	err = tx.QueryRow(`
		SELECT wallet_balance FROM users WHERE id = $1
	`, userID).Scan(&walletBalance)
	if err != nil {
		return err
	}

	if walletBalance < amount {
		return errors.New("insufficient balance")
	}

	_, err = tx.Exec(`
		UPDATE users SET wallet_balance = wallet_balance - $1 WHERE id = $2
	`, amount, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE lotto_numbers SET is_sold = true, sold_to = $1 WHERE id = $2 AND is_sold = false
	`, userID, lottoNumberID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions (user_id, lotto_number_id, transaction_type, amount)
		VALUES ($1, $2, 'purchase', $3)
	`, userID, lottoNumberID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// DrawLottoNumbers performs a lotto draw
func DrawLottoNumbers() ([]string, error) {
	rows, err := db.Conn.Query(`
		SELECT number FROM lotto_numbers WHERE is_sold = true
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []string
	for rows.Next() {
		var num string
		if err := rows.Scan(&num); err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if len(numbers) < 5 {
		return nil, errors.New("not enough sold lotto numbers to draw")
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
	return numbers[:5], nil
}

// ClaimWinnings processes claiming of winnings
func ClaimWinnings(userID int) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rows, err := tx.Query(`
		SELECT l.id, p.prize_amount
		FROM lotto_numbers l
		JOIN prizes p ON l.id = p.lotto_number_id
		WHERE l.sold_to = $1 AND p.claimed = false
	`, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var totalWinnings int
	var lottoID int
	for rows.Next() {
		var prizeAmount int
		if err := rows.Scan(&lottoID, &prizeAmount); err != nil {
			return err
		}
		totalWinnings += prizeAmount

		_, err := tx.Exec(`
			UPDATE prizes SET claimed = true WHERE lotto_number_id = $1
		`, lottoID)
		if err != nil {
			return err
		}
	}

	if totalWinnings > 0 {
		_, err = tx.Exec(`
			UPDATE users SET wallet_balance = wallet_balance + $1 WHERE id = $2
		`, totalWinnings, userID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
