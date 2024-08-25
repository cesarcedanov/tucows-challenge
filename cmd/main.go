package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tucows-challenge/model"
	"tucows-challenge/server"
)

func main() {
	fmt.Printf("Server started at %s", time.Now())
	fmt.Printf("Show Menu \n %+v", model.Menu)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"healthy-check": "tucows-coffee"})
	})
	router.GET(urLWithPrefix("menu"), server.GetMenu)
	router.GET(urLWithPrefix("order/all"), server.GetAllOrders)
	router.GET(urLWithPrefix("order/:id"), server.GetOrder)
	router.POST(urLWithPrefix("order"), server.CreateOrder)
	router.PUT(urLWithPrefix("order/:id"), server.UpdateOrder)
	router.PATCH(urLWithPrefix("order/:id/confirm"), server.ConfirmOrder)
	router.DELETE(urLWithPrefix("order/:id/cancel"), server.CancelOrder)

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Printf("Server stopped at %s", time.Now())
		fmt.Printf("Internal Error: %s", err)
	}
}

func urLWithPrefix(url string) string {
	return fmt.Sprintf("/tucows-coffee/%s", url)
}
