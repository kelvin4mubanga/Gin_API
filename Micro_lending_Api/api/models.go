package api

import "time"

type Loan struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Borrower  string    `json:"borrower"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}