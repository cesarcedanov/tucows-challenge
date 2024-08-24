package store

import (
	"time"
	"tucows-challenge/model"
)

var Orders = map[int]model.Order{
	1: model.Order{
		ID:         1,
		ClientName: "Cesar",
		Status:     model.OrderStatus_InProgress,
		Products: []model.Product{
			model.Espresso,
			model.Americano,
			model.Americano,
			model.TuCowsMilk,
		},
		TotalPrice: 12, // This should be Calculated on Runtime
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		UpdatedBy:  "Employee #1",
	},
	2: model.Order{
		ID:         2,
		ClientName: "Jon Doe",
		Status:     model.OrderStatus_InProgress,
		Products: []model.Product{
			model.Espresso,
			model.TuCowsMilk,
		},
		TotalPrice: 7, // This should be Calculated on Runtime
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		UpdatedBy:  "Employee #2",
	},
}
