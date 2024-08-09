package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	PasswordHash  string    `json:"-"`
	Role          string    `json:"role"`
	WalletBalance int       `json:"wallet_balance"`
	CreatedAt     time.Time `json:"created_at"`
}
