package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
	"tucows-challenge/api/server"
	"tucows-challenge/api/service"
	"tucows-challenge/api/store"
)

func main() {
	log.Printf("Server started at %s", time.Now())
	db := store.InitDB()
	kitchen := service.NewKitchen(5, 100, db)

	handler := server.OrderHandler{
		Kitchen: kitchen,
		StoreDB: db,
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"healthy-check": "tucows-coffee"})
	})

	// Login
	router.POST(urLWithPrefix("login"), server.Login)

	// Non-Auth
	router.GET(urLWithPrefix("menu"), handler.GetMenu)

	// Auth
	router.GET(urLWithPrefix("order/all"), server.MiddlewareAuth(), handler.GetAllOrders)
	router.GET(urLWithPrefix("order/:id"), server.MiddlewareAuth(), handler.GetOrder)
	router.POST(urLWithPrefix("order"), server.MiddlewareAuth(), handler.CreateOrder)
	router.PUT(urLWithPrefix("order/:id"), server.MiddlewareAuth(), handler.UpdateOrder)
	router.DELETE(urLWithPrefix("order/:id/cancel"), server.MiddlewareAuth(), handler.ChangeOrderStatus)
	router.PATCH(urLWithPrefix("order/:id/confirm"), server.MiddlewareAuth(), handler.ChangeOrderStatus)
	router.PATCH(urLWithPrefix("order/confirm/all"), server.MiddlewareAuth(), handler.ConfirmPreOrders)

	certFile := "./cert/server.crt"
	keyFile := "./cert/server.key"
	if os.Getenv("IS_DOCKERIZED") == "true" {
		certFile = "/app/cmd/cert/server.crt"
		keyFile = "/app/cmd/cert/server.key"
	}

	if err := router.RunTLS(":8080", certFile, keyFile); err != nil {
		log.Printf("Server stopped at %s", time.Now())
		log.Printf("Internal Error: %s", err)
	}
}

func urLWithPrefix(url string) string {
	return fmt.Sprintf("/tucows-coffee/%s", url)
}
