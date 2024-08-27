package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"tucows-challenge/api/model"
	"tucows-challenge/api/service"
)

type OrderHandler struct {
	Kitchen service.KitchenService
	StoreDB *gorm.DB
}

const kEmployee = "Tester"

func (handler *OrderHandler) GetMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Menu)
}

func (handler *OrderHandler) GetAllOrders(c *gin.Context) {
	rows := []model.Order{}
	handler.StoreDB.Model(&model.Order{}).Find(&rows)

	resp := []model.OrderResponse{}
	for _, order := range rows {
		resp = append(resp, model.HumanizeOrder(&order))
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
	newOrder.Status = model.OrderStatus_PreOrder
	model.CalculateOrderPrice(newOrder)
	newOrder.CreatedAt = time.Now()
	newOrder.UpdatedAt = time.Now()
	newOrder.UpdatedBy = kEmployee

	handler.StoreDB.Model(&model.Order{}).Create(newOrder)

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
	handler.StoreDB.Save(&existingOrder)

	c.IndentedJSON(http.StatusOK, existingOrder)
}

func (handler *OrderHandler) ChangeOrderStatus(c *gin.Context) {
	order := handler.getOrderByID(c)
	if order.Status != model.OrderStatus_PreOrder {
		c.IndentedJSON(http.StatusForbidden, gin.H{"errorMsg": fmt.Sprintf("Orders %v is not pre-order", order.ID)})
	}
	order.UpdatedAt = time.Now()
	order.UpdatedBy = kEmployee
	switch c.Request.Method {
	case http.MethodPatch:
		order.Status = model.OrderStatus_Confirmed
		handler.StoreDB.Save(order)
		handler.Kitchen.AddConfirmedOrder(order)
	case http.MethodDelete:
		order.Status = model.OrderStatus_Canceled
		handler.StoreDB.Delete(order)
	}
	c.IndentedJSON(http.StatusOK, order)
}

func (handler *OrderHandler) getOrderByID(c *gin.Context) *model.Order {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": "Orders ID should be an integer"})
		return nil
	}
	order := &model.Order{}
	if result := handler.StoreDB.Find(order, orderID); result.Error != nil {
		if result.RowsAffected == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"errorMsg": "Order not found"})
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": result.Error.Error()})
		return nil
	}

	return order
}
