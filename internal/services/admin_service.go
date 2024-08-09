package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"lottery-api/internal/db"
	"lottery-api/internal/models"
	"math/rand"
	"os"
	"time"

	"github.com/lib/pq"
)

// Prize amounts
var prizeAmounts = map[int]int{
	1: 10000, // 1st place
	2: 8000,  // 2nd place
	3: 6000,  // 3rd place
	4: 4000,  // 4th place
	5: 2000,  // 5th place
}

// StoreDrawnNumbers stores the drawn lotto numbers in the database
func StoreDrawnNumbers(numbers []models.LottoNumber) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert draw record
	drawDate := time.Now()
	var drawID int
	err = tx.QueryRow(`
		INSERT INTO draws (draw_date, winning_numbers)
		VALUES ($1, $2) RETURNING id
	`, drawDate, pq.Array(GetNumbersString(numbers))).Scan(&drawID)
	if err != nil {
		return err
	}

	// Insert prizes for the drawn numbers based on their rank
	for rank, num := range numbers {
		if prizeAmount, ok := prizeAmounts[rank]; ok {
			_, err = tx.Exec(`
				INSERT INTO prizes (draw_id, lotto_number_id, prize_amount)
				VALUES ($1, $2, $3)
			`, drawID, num.ID, prizeAmount)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

// GetPrizeAmount determines the prize amount based on the lotto number's rank
func GetPrizeAmount(rank int) int {
	return prizeAmounts[rank]
}

// GetNumbersString converts a slice of LottoNumber to a string slice for storage
func GetNumbersString(numbers []models.LottoNumber) []string {
	var result []string
	for _, num := range numbers {
		result = append(result, num.Number)
	}
	return result
}

// GenerateRandomNumbers selects a specified number of random lotto numbers from  numbers
func GenerateRandomNumbers(count int) ([]models.LottoNumber, error) {
	Numbers, err := GetAllLottoNumbers()
	if err != nil {
		return nil, err
	}

	if len(Numbers) < count {
		return nil, errors.New("not enough  lotto numbers to draw")
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle  numbers
	rand.Shuffle(len(Numbers), func(i, j int) { Numbers[i], Numbers[j] = Numbers[j], Numbers[i] })

	// Select the first 'count' numbers
	return Numbers[:count], nil
}

// ResetSystem clears all data except for admin's data
func ResetSystem() error {
	// Read SQL file
	file, err := os.Open("init_db.sql")
	if err != nil {
		return fmt.Errorf("failed to open SQL file: %w", err)
	}
	defer file.Close()

	sqlBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	sqlStatements := string(sqlBytes)

	// Execute SQL statements
	tx, err := db.Conn.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(sqlStatements)
	if err != nil {
		return fmt.Errorf("failed to execute SQL statements: %w", err)
	}

	return tx.Commit()
}
