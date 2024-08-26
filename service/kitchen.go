package service

import (
	"sync"
	"tucows-challenge/model"
)

type KitchenService interface {
	AddConfirmedOrder(order *model.Order)
	GetOrders() map[int]*model.Order
}

type Kitchen struct {
	Workers int // Total Workers
	Orders  map[int]*model.Order
	Queue   chan *model.Order // Channel
	mutex   sync.Mutex
}

func NewKitchen(workers int, concurrentOrders int, orders map[int]*model.Order) *Kitchen {
	kitchen := &Kitchen{
		Workers: workers,
		Queue:   make(chan *model.Order, concurrentOrders),
		Orders:  orders,
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
		order.Prepare(workerID)
		k.mutex.Lock()
		k.Orders[order.ID] = order
		k.mutex.Unlock()
	}
}

func (k *Kitchen) AddConfirmedOrder(order *model.Order) {
	k.mutex.Lock()
	// Added the new Confirmed Orders into the Kitchen
	k.Orders[order.ID] = order
	k.mutex.Unlock()
	// Push it into the Queue
	k.Queue <- order
}

func (k *Kitchen) GetOrders() map[int]*model.Order {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	return k.Orders
}

func (k *Kitchen) Close() {
	// We need to Close the channel in case we close the Kitchen.
	close(k.Queue)
}
