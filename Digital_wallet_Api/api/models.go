package api

import "time"

type Wallet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Owner     string    `json:"owner"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}