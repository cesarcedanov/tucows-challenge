package main

import "fmt"

type Prettier struct{}

func (p Prettier) MenuProducts(products []Product) string {
	str := fmt.Sprintf("\nTuCows MENU")
	str += "\n"
	str += p.ProductsDetails(products)
	return str
}

func (p Prettier) OrdersDetails(orders []OrderResponse, showProducts bool) string {
	str := "\n"
	for _, order := range orders {
		str += fmt.Sprintf("\n-> Order #%v for Client (%s) is on Status (%s)", order.ID, order.ClientName, order.Status)
		if showProducts {
			str += p.ProductsDetails(order.Products)
		}
		str += "\n"
	}

	return str
}

func (p Prettier) ProductsDetails(products []Product) string {
	str := fmt.Sprintf("")
	for i, product := range products {
		str += fmt.Sprintf("\n%v) %s CAD$%v", i+1, product.Title, product.Price)
	}
	return str
}
