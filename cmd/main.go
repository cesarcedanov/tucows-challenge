package main

import (
	"fmt"
	"net/http"
	"time"
	"tucows-challenge/model"
	"tucows-challenge/server"
)

func main() {
	fmt.Printf("Server started at %s", time.Now())
	fmt.Printf("Show Menu \n %+v", model.Menu)
	http.HandleFunc(urLWithPrefix("menu"), server.GetMenu)
	http.HandleFunc(urLWithPrefix("order/all"), server.GetAllOrders)
	http.HandleFunc(urLWithPrefix("order/1"), server.GetOrder)
	http.ListenAndServe(":8080", nil)
}

func urLWithPrefix(url string) string {
	return fmt.Sprintf("/tucows-coffee/%s", url)
}
