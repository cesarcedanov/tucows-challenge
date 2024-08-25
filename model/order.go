package model

import "time"

type Order struct {
	ID         int        `json:"id"`
	ClientName string     `json:"client_name"`
	Status     string     `json:"status"`
	Products   []int      `json:"products"`
	Price      OrderPrice `json:"price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UpdatedBy  string     `json:"updated_by"`
}

type OrderRequest struct {
	ClientName string `json:"client_name"`
	//Status     string     `json:"status"`
	Products  []int      `json:"products"`
	Price     OrderPrice `json:"price"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type OrderResponse struct {
	ID         int        `json:"id"`
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
