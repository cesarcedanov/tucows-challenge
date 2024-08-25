package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tucows-challenge/model"
	"tucows-challenge/store"
	"tucows-challenge/utils"
)

//	type OrderServer struct {
//		// dbStore
//		// workerService
//	}
const kEmployee = "Tester"

func GetAllOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.Orders)
}

func GetOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(400)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": "Order ID should be an integer"})
		return
	}
	order, found := store.Orders[orderID]
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"errorMsg": "Order not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, order)

}

func CreateOrder(c *gin.Context) {
	newOrder := &model.Order{}
	if err := c.BindJSON(newOrder); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}
	nextID := store.GetNextID()
	newOrder.ID = nextID
	newOrder.Status = model.OrderStatus_PreOrder
	utils.CalculateOrderPrice(newOrder)
	newOrder.CreatedAt = time.Now()
	newOrder.UpdatedAt = time.Now()
	newOrder.UpdatedBy = kEmployee

	store.Orders[nextID] = newOrder

	c.IndentedJSON(http.StatusCreated, newOrder)
}

func UpdateOrder(c *gin.Context) {
	existingOrder := getOrderByID(c)
	updateOrder := &model.Order{}
	if err := c.BindJSON(updateOrder); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}
	if len(updateOrder.Products) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": fmt.Sprintf("Products should not be empty or Cancel the Order %s", updateOrder.ID)})
	}

	existingOrder.ClientName = updateOrder.ClientName
	existingOrder.Products = updateOrder.Products
	utils.CalculateOrderPrice(existingOrder)
	existingOrder.Price = updateOrder.Price
	existingOrder.UpdatedAt = time.Now()
	existingOrder.UpdatedBy = kEmployee

	c.IndentedJSON(http.StatusOK, existingOrder)
}

// Don't repeat yourself - CancelOrder
func ConfirmOrder(c *gin.Context) {
	order := getOrderByID(c)
	if order.Status != model.OrderStatus_PreOrder {
		c.IndentedJSON(http.StatusForbidden, gin.H{"errorMsg": fmt.Sprintf("Order %s is not pre-order", order.ID)})
	}
	order.Status = model.OrderStatus_Confirmed
	order.UpdatedAt = time.Now()
	order.UpdatedBy = kEmployee
	store.Orders[order.ID] = order
	c.IndentedJSON(http.StatusAccepted, order)
}

// Don't repeat yourself - ConfirmOrder
func CancelOrder(c *gin.Context) {
	order := getOrderByID(c)
	if order.Status != model.OrderStatus_PreOrder {
		c.IndentedJSON(http.StatusForbidden, gin.H{"errorMsg": fmt.Sprintf("Order %s is not pre-order", order.ID)})
	}
	order.Status = model.OrderStatus_Canceled
	order.UpdatedAt = time.Now()
	order.UpdatedBy = kEmployee
	store.Orders[order.ID] = order
	c.IndentedJSON(http.StatusOK, order)
}

func getOrderByID(c *gin.Context) *model.Order {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errorMsg": "Order ID should be an integer"})
		return nil
	}
	order, found := store.Orders[orderID]
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"errorMsg": "Order not found"})
		return nil
	}
	return order
}
