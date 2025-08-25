package api

import "time"

type Reservation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Vehicle   string    `json:"vehicle"`
	Customer  string    `json:"customer"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}