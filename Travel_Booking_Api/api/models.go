package api

import "time"

type Booking struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Customer    string    `json:"customer"`
	Destination string    `json:"destination"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}