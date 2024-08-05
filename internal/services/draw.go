package services

import (
	"lottery-api/internal/db"
	"lottery-api/internal/models"
)

func GetDrawResults() ([]models.LotteryResult, error) {
	rows, err := db.DB.Query("SELECT * FROM lottery_results ORDER BY draw_date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.LotteryResult
	for rows.Next() {
		var result models.LotteryResult
		if err := rows.Scan(&result.DrawID, &result.DrawDate, &result.Number1, &result.Number2, &result.Number3, &result.Number4, &result.Number5); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
