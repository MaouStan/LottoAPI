package models

import "time"

// LottoNumber represents a lotto number
type LottoNumber struct {
	ID        int       `json:"id"`
	Number    string    `json:"number"`
	IsSold    bool      `json:"is_sold"`
	SoldTo    *int      `json:"sold_to"` // Nullable
	CreatedAt time.Time `json:"created_at"`
}

// Transaction represents a transaction record
type Transaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	LottoNumberID   int       `json:"lotto_number_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          int       `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}

// Draw represents a lotto draw
type Draw struct {
	ID             int       `json:"id"`
	DrawDate       time.Time `json:"draw_date"`
	WinningNumbers []string  `json:"winning_numbers"`
}

// Prize represents a prize record
type Prize struct {
	ID            int  `json:"id"`
	DrawID        int  `json:"draw_id"`
	LottoNumberID int  `json:"lotto_number_id"`
	PrizeAmount   int  `json:"prize_amount"`
	Claimed       bool `json:"claimed"`
}
