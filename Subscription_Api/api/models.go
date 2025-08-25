package api

import "time"

type Subscription struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	User      string    `json:"user"`
	Plan      string    `json:"plan"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}