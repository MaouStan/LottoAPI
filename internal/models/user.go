package models

// RegisterRequest represents the payload for user registration.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequest represents the payload for user login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User represents a general user in the system.
type User struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	IsMember     bool   `json:"is_member"`
}

// Member represents a member in the system, which includes additional wallet information.
type Member struct {
	MemberID      int     `json:"member_id"`
	WalletBalance float64 `json:"wallet_balance"`
}
