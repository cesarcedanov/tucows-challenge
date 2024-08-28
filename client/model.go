package main

import (
	"time"
)

type Product struct {
	Title       string           `json:"title"`
	Price       float32          `json:"price"`
	Ingredients map[string]int32 `json:"ingredients"`
}

type OrderRequest struct {
	ID         uint       `json:"id"`
	ClientName string     `json:"client_name"`
	Status     string     `json:"status"`
	Products   []int      `json:"products"`
	Price      OrderPrice `json:"price"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type OrderResponse struct {
	ID         uint       `json:"id"`
	ClientName string     `json:"client_name"`
	Status     string     `json:"status"`
	Products   []Product  `json:"products"`
	Price      OrderPrice `json:"price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UpdatedBy  string     `json:"updated_by"`
}

type OrderPrice struct {
	FinalPrice float32 `json:"final_price"`
	AutoPrice  bool    `json:"auto_price"`
}
