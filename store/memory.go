package store

import (
	"time"
	"tucows-challenge/model"
)

// NextIDSequence will increase based on InitOrders created
var NextIDSequence = len(InitOrders)

func GetNextID() int {
	NextIDSequence++
	return NextIDSequence
}

var InitOrders = map[int]*model.Order{
	1: &model.Order{
		ID:         1,
		ClientName: "Cesar",
		Status:     model.OrderStatus_InProgress,
		Products: []int{
			model.ProductID_Espresso,
			model.ProductID_Americano,
			model.ProductID_Americano,
			model.ProductID_TuCowsMilk,
		},
		Price: model.OrderPrice{
			FinalPrice: 12,
			AutoPrice:  true,
		}, // This should be Calculated on Runtime
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UpdatedBy: "Employee #1",
	},
	2: &model.Order{
		ID:         2,
		ClientName: "Jon Doe",
		Status:     model.OrderStatus_InProgress,
		Products: []int{
			model.ProductID_Espresso,
			model.ProductID_TuCowsMilk,
		},
		Price: model.OrderPrice{
			AutoPrice:  false,
			FinalPrice: 20,
		}, // This should be Calculated on Runtime
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UpdatedBy: "Employee #2",
	},
}
