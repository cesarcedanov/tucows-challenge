package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tucows-challenge/model"
	"tucows-challenge/service"
	"tucows-challenge/store"
)

type OrderHandler struct {
	Kitchen service.KitchenService
}

const kEmployee = "Tester"

func (handler *OrderHandler) GetMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Menu)
}

func (handler *OrderHandler) GetAllOrders(c *gin.Context) {
	resp := []model.OrderResponse{}
	for _, order := range handler.Kitchen.GetOrders() {
		resp = append(resp, model.HumanizeOrder(order))
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func (handler *OrderHandler) GetOrder(c *gin.Context) {
	order := handler.getOrderByID(c)
	c.IndentedJSON(http.StatusOK, model.HumanizeOrder(order))

}

func (handler *OrderHandler) CreateOrder(c *gin.Context) {
	newOrder := &model.Order{}
	if err := c.BindJSON(newOrder); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}
	nextID := store.GetNextID()
	newOrder.ID = nextID
	newOrder.Status = model.OrderStatus_PreOrder
	model.CalculateOrderPrice(newOrder)
	newOrder.CreatedAt = time.Now()
	newOrder.UpdatedAt = time.Now()
	newOrder.UpdatedBy = kEmployee

	handler.Kitchen.GetOrders()[nextID] = newOrder

	c.IndentedJSON(http.StatusCreated, newOrder)
}

func (handler *OrderHandler) UpdateOrder(c *gin.Context) {
	existingOrder := handler.getOrderByID(c)
	updateOrder := &model.Order{}
	if err := c.BindJSON(updateOrder); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}
	if len(updateOrder.Products) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": fmt.Sprintf("Products should not be empty or Cancel the Orders %v", updateOrder.ID)})
	}

	existingOrder.ClientName = updateOrder.ClientName
	existingOrder.Products = updateOrder.Products
	existingOrder.Price = updateOrder.Price
	model.CalculateOrderPrice(existingOrder)
	existingOrder.UpdatedAt = time.Now()
	existingOrder.UpdatedBy = kEmployee

	c.IndentedJSON(http.StatusOK, existingOrder)
}

func (handler *OrderHandler) ChangeOrderStatus(c *gin.Context) {
	order := handler.getOrderByID(c)
	if order.Status != model.OrderStatus_PreOrder {
		c.IndentedJSON(http.StatusForbidden, gin.H{"errorMsg": fmt.Sprintf("Orders %v is not pre-order", order.ID)})
	}
	switch c.Request.Method {
	case http.MethodPatch:
		order.Status = model.OrderStatus_Confirmed
		handler.Kitchen.AddConfirmedOrder(order)
	case http.MethodDelete:
		order.Status = model.OrderStatus_Canceled
		delete(handler.Kitchen.GetOrders(), order.ID)
	}
	order.UpdatedAt = time.Now()
	order.UpdatedBy = kEmployee
	handler.Kitchen.GetOrders()[order.ID] = order
	c.IndentedJSON(http.StatusOK, order)
}

func (handler *OrderHandler) getOrderByID(c *gin.Context) *model.Order {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": "Orders ID should be an integer"})
		return nil
	}
	kitchenOrders := handler.Kitchen.GetOrders()
	order, found := kitchenOrders[orderID]
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"errorMsg": "Orders not found"})
		return nil
	}
	return order
}
