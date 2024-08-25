package utils

import "tucows-challenge/model"

func productµMapper(ID int) model.Product {
	switch ID {
	case model.ProductID_Espresso:
		return model.Espresso
	case model.ProductID_Americano:
		return model.Americano
	case model.ProductID_Cappuccino:
		return model.Cappuccino
	case model.ProductID_Latte:
		return model.Latte
	case model.ProductID_TuCowsMilk:
		return model.TuCowsMilk
	default:
		return model.Product{
			Title: "Invalid Product",
		}
	}
}

func CalculateOrderPrice(order *model.Order) {
	var finalPrice float32
	if order.Price.AutoPrice {
		for _, productID := range order.Products {
			product := productµMapper(productID)
			finalPrice += product.Price
		}
		order.Price.FinalPrice = finalPrice
	}
}
