package store

import (
	"time"
	"tucows-challenge/api/model"
)

var InitOrders = []*model.Order{
	1: &model.Order{
		ID:         1,
		ClientName: "Cesar",
		Status:     model.OrderStatus_PreOrder,
		Products: []int64{
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
		Status:     model.OrderStatus_PreOrder,
		Products: []int64{
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
