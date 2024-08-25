package model

import (
	"fmt"
	"time"
)

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

func (o *Order) Prepare(workerID int) {
	o.Status = OrderStatus_InProgress
	o.UpdatedAt = time.Now()
	o.UpdatedBy = fmt.Sprintf("Worker#%d", workerID)

	for _, p := range o.Products {
		product := productMapper(p)
		fmt.Printf("Worker %v preparing a %s for Orders ID %v\n", workerID, product.Title, o.ID)
		if effort, found := product.Ingredients[Coffee]; found {
			fmt.Printf("Worker %v: Grinding coffee beans and Brewing it\n", workerID)
			time.Sleep(time.Microsecond * time.Duration(effort))
		}
		if effort, found := product.Ingredients[Milk]; found {
			fmt.Printf("Worker %v: Steaming Milk\n", workerID)
			time.Sleep(time.Microsecond * time.Duration(effort))
		}
		if effort, found := product.Ingredients[Water]; found {
			fmt.Printf("Worker %v: Pouring Water\n", workerID)
			time.Sleep(time.Microsecond * time.Duration(effort))
		}
	}
	o.Status = OrderStatus_Finished
}

func CalculateOrderPrice(order *Order) {
	var finalPrice float32
	if order.Price.AutoPrice {
		for _, productID := range order.Products {
			product := productMapper(productID)
			finalPrice += product.Price
		}
		order.Price.FinalPrice = finalPrice
	}
	// Final Price can be Zero for FREE Coffee
}

func HumanizeOrder(o *Order) OrderResponse {
	products := make([]Product, 0)
	for _, productID := range o.Products {
		products = append(products, productMapper(productID))
	}

	return OrderResponse{
		ID:         o.ID,
		ClientName: o.ClientName,
		Status:     o.Status,
		Products:   products,
		Price:      o.Price,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
		UpdatedBy:  o.UpdatedBy,
	}
}
