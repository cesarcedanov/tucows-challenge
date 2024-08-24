package server

import (
	"fmt"
	"net/http"
	"tucows-challenge/store"
)

//type OrderServer struct {
//	// dbStore
//	// workerService
//}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Order")
}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Confirm Order")
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: \n\n %+v", r)
	fmt.Fprintf(w, "Get Order")
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get All Orders: %+v", store.Orders)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Order")
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Cancel Order")
}
