package model

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Order struct {
	ID         uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientName string        `json:"client_name" gorm:"column:client_name;not null"`
	Status     string        `json:"status" gorm:"column:status;not null"`
	Products   pq.Int64Array `json:"products" gorm:"type:integer[]"`
	Price      OrderPrice    `json:"price" gorm:"embedded"`
	CreatedAt  time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time     `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy  string        `json:"updated_by" gorm:"column:updated_by"`
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

func (o *Order) Prepare(workerID int, db *gorm.DB) {
	o.Status = OrderStatus_InProgress
	o.UpdatedAt = time.Now()
	o.UpdatedBy = fmt.Sprintf("Worker#%d", workerID)
	db.Save(&o)
	for _, p := range o.Products {
		product := productMapper(p)
		log.Printf("Worker %v preparing a %s for Orders ID %v\n", workerID, product.Title, o.ID)
		if effort, found := product.Ingredients[Coffee]; found {
			log.Printf("Worker %v: Grinding coffee beans and Brewing it\n", workerID)
			time.Sleep(time.Microsecond * time.Duration(effort))
		}
		if effort, found := product.Ingredients[Milk]; found {
			log.Printf("Worker %v: Steaming Milk\n", workerID)
			time.Sleep(time.Microsecond * time.Duration(effort))
		}
		if effort, found := product.Ingredients[Water]; found {
			log.Printf("Worker %v: Pouring Water\n", workerID)
			time.Sleep(time.Microsecond * time.Duration(effort))
		}
	}
	o.Status = OrderStatus_Finished
	o.UpdatedAt = time.Now()
	log.Printf("Worker %v finished Order %v\n", workerID, o.ID)
	db.Save(&o)
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
