package entity

import "time"

// Expense represents a single expense
type Expense struct {
	ID          int       `json:"id"`                 // Unique ID of the expense
	Description string    `json:"description"`        // Description of the expense
	Amount      float64   `json:"amount"`             // Amount spent
	Date        time.Time `json:"date"`               // Date of the expense
	Category    string    `json:"category,omitempty"` // Optional category
}
