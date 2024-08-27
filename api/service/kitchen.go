package service

import (
	"gorm.io/gorm"
	"tucows-challenge/api/model"
)

type KitchenService interface {
	AddConfirmedOrder(order *model.Order)
}

type Kitchen struct {
	Workers int               // Total Workers
	Queue   chan *model.Order // Channel
	StoreDB *gorm.DB
	//mutex   sync.Mutex
}

func NewKitchen(workers int, concurrentOrders int, db *gorm.DB) *Kitchen {
	kitchen := &Kitchen{
		Workers: workers,
		Queue:   make(chan *model.Order, concurrentOrders),
		StoreDB: db,
	}
	kitchen.run()
	return kitchen
}

func (k *Kitchen) run() {
	for i := 1; i <= k.Workers; i++ {
		go k.processOrder(i)
	}
}

func (k *Kitchen) processOrder(workerID int) {
	for order := range k.Queue {
		order.Prepare(workerID, k.StoreDB)
	}
}

func (k *Kitchen) AddConfirmedOrder(order *model.Order) {
	// Push it into the Queue
	k.Queue <- order
}

func (k *Kitchen) Close() {
	// We need to Close the channel in case we close the Kitchen.
	close(k.Queue)
}
