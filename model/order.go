package model

import "time"

type Order struct {
	ID         int       `json:"id"`
	ClientName string    `json:"client_name"`
	Status     string    `json:"status"`
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
}
