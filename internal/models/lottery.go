package models

// LotteryNumber represents a lottery number entry.
type LotteryNumber struct {
	NumberID    string `json:"number_id"`    // e.g., "LT-000001"
	NumberValue string `json:"number_value"` // e.g., "000001"
}

// DrawNumber represents a lottery number available for a specific draw.
type DrawNumber struct {
	DrawID    int    `json:"draw_id"`
	NumberID  string `json:"number_id"` // Changed to string to match database schema
	Available bool   `json:"available"`
}

// PurchaseRequest represents a request to purchase a lottery ticket.
type PurchaseRequest struct {
	MemberID int     `json:"member_id"`
	NumberID string  `json:"number_id"` // Changed to string to match database schema
	Amount   float64 `json:"amount"`
}

// TransferRequest represents a request to transfer a lottery ticket from one member to another.
type TransferRequest struct {
	FromMemberID int    `json:"from_member_id"`
	ToMemberID   int    `json:"to_member_id"`
	NumberID     string `json:"number_id"` // Changed to string to match database schema
}
