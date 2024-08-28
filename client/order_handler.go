package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	url = "https://localhost:8080/tucows-coffee/"
)

var client = &http.Client{}
var prettier = Prettier{}

func initClient() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client = &http.Client{Transport: tr}

}

// fetch Menu from the API Server
func fetchMenu() {
	resp, err := client.Get(url + "menu")
	if err != nil {
		log.Fatal("Error fetching menu data:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	products := []Product{}
	jsonErr := json.Unmarshal(body, &products)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.MenuProducts(products))
}

// fetch All Orders from the API Server
func fetchAllOrders() {
	req, err := http.NewRequest(http.MethodGet, url+"order/all", nil)
	if err != nil {
		log.Fatalf("Failed Get Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error fetching all Orders:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	orders := []OrderResponse{}
	jsonErr := json.Unmarshal(body, &orders)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails(orders, false))
}

// fetch specific Order by ID from the API Server
func fetchOrderByID(orderID string) {
	req, err := http.NewRequest(http.MethodGet, url+"order/"+orderID, nil)
	if err != nil {
		log.Fatalf("Failed Get Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error fetching Order:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	order := OrderResponse{}
	jsonErr := json.Unmarshal(body, &order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails([]OrderResponse{order}, true))
}

// populate the data for a Request
func populateOrderRequest(id string) *OrderRequest {
	var orderID int
	var err error
	if id != "" {
		orderID, err = strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid Order ID:", orderID)
			fmt.Println("Try again...")
			return nil
		}
	}

	clientName := readInput("Enter your Client's name: ")
	products := readInput("Enter your Products separated by comma, (ex: 1,2,3,4)  :")
	productsIDs := strings.Split(strings.Replace(products, " ", "", -1), ",")
	productsArray := []int{}
	for _, productID := range productsIDs {
		if n, err := strconv.Atoi(productID); err == nil {
			productsArray = append(productsArray, n)
			continue
		}
		fmt.Println("Invalid Product ID:", productID)
		fmt.Println("Try again...")
		return nil
	}

	autoPrice := true
	var finalPrice float32
	isAutoPrice := readInput("Price is auto Calculated, but you can alter the price manually. Do you want set price manually it? Y / N:")
	if isAutoPrice == "Y" || isAutoPrice == "y" {
		autoPrice = false
		manualPrice := readInput("Enter the final price: ")
		if n, err := strconv.Atoi(manualPrice); err == nil {
			finalPrice = float32(n)
		} else {
			fmt.Println("Invalid Decimal Value:", manualPrice)
			fmt.Println("Try again...")
			return nil
		}
	}

	return &OrderRequest{
		ID:         uint(orderID),
		ClientName: clientName,
		Products:   productsArray,
		Price: OrderPrice{
			FinalPrice: finalPrice,
			AutoPrice:  autoPrice,
		},
	}
}

func createOrder(payload OrderRequest) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling data:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, url+"order", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed Post Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error creating Order:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	order := OrderResponse{}
	jsonErr := json.Unmarshal(body, &order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails([]OrderResponse{order}, true))

}

func updateOrder(payload OrderRequest) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling data:", err)
		return
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%sorder/%v", url, payload.ID), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed Put Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error updating Order:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	order := OrderResponse{}
	jsonErr := json.Unmarshal(body, &order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails([]OrderResponse{order}, true))

}

func confirmOrder(orderID string) {
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprint(url+"order/"+orderID+"/confirm"), nil)
	if err != nil {
		log.Fatalf("Failed Patch Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error confirming Order:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	order := OrderResponse{}
	jsonErr := json.Unmarshal(body, &order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails([]OrderResponse{order}, true))
}

func cancelOrder(orderID string) {

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprint(url+"order/"+orderID+"/cancel"), nil)
	if err != nil {
		log.Fatalf("Failed Delete Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error deleting Order:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	order := OrderResponse{}
	jsonErr := json.Unmarshal(body, &order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails([]OrderResponse{order}, true))
}

func confirmAllPreOrder() {
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprint(url+"order/confirm/all"), nil)
	if err != nil {
		log.Fatalf("Failed Patch Request: %v", err)
	}
	req.Header.Set("Authorization", jwtToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error confirming All Orders", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	orders := []OrderResponse{}
	jsonErr := json.Unmarshal(body, &orders)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(prettier.OrdersDetails(orders, false))
}
