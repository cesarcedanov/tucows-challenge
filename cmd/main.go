package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tucows-challenge/server"
	"tucows-challenge/service"
	"tucows-challenge/store"
)

func main() {
	fmt.Printf("Server started at %s", time.Now())
	// Init Kitchen Queue

	kitchen := service.NewKitchen(5, 100, store.InitOrders)

	handler := server.OrderHandler{
		kitchen,
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"healthy-check": "tucows-coffee"})
	})
	router.GET(urLWithPrefix("menu"), handler.GetMenu)
	router.GET(urLWithPrefix("order/all"), handler.GetAllOrders)
	router.GET(urLWithPrefix("order/:id"), handler.GetOrder)
	router.POST(urLWithPrefix("order"), handler.CreateOrder)
	router.PUT(urLWithPrefix("order/:id"), handler.UpdateOrder)
	router.PATCH(urLWithPrefix("order/:id/confirm"), handler.ChangeOrderStatus)
	router.DELETE(urLWithPrefix("order/:id/cancel"), handler.ChangeOrderStatus)

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Printf("Server stopped at %s", time.Now())
		fmt.Printf("Internal Error: %s", err)
	}
}

func urLWithPrefix(url string) string {
	return fmt.Sprintf("/tucows-coffee/%s", url)
}
