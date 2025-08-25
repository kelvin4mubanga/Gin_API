package api

import "time"

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}