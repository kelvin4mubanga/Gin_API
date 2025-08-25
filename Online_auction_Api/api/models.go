package api

import "time"

type Auction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Item        string    `json:"item"`
	StartingBid float64   `json:"starting_bid"`
	Status      string    `json:"status"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}