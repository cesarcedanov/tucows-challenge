package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var currentMenu int

const (
	PublicMenu = iota
	EmployeeMenu
)

// Show the initial Menu
func displayGuestMenu() string {
	fmt.Println("Please select an option:")
	fmt.Println("1. Login as Employee")
	fmt.Println("2. Show Products Menu")
	fmt.Println("3. Exit")
	choice := readInput("Enter your choice: ")
	if !executeOperation(choice) {
		os.Exit(1)

	}
	return choice
}

// Function to execute the selected operation
func executeOperation(choice string) bool {
	switch choice {
	case "1":
		login()
		if jwtToken != "" {
			fmt.Printf("\nWelcome '%s' - Sesion started at %s", employeeUsername, time.Now().String())
			currentMenu = EmployeeMenu
			for currentMenu == EmployeeMenu {
				displayEmployeeMenu()
			}
		} else {
			waitUntilPress()
		}

	case "2":
		fetchMenu()
	case "3":
		// Fetch Menu
		fmt.Println("Exiting TuCows Coffee - Bye!")
		return false
	default:
		fmt.Println("Invalid option, please try again.")
	}
	return true
}

// Show the initial Menu
func displayEmployeeMenu() {
	fmt.Printf("\n\nCurrent Employee '%s'\n", employeeUsername)
	fmt.Println("0. Show Products Menu")
	fmt.Println("1. Check ALL Orders")
	fmt.Println("2. Check Order by ID")
	fmt.Println("3. Create Order")
	fmt.Println("4. Update Order")
	fmt.Println("5. Cancel Order")
	fmt.Println("6. Confirm Order")
	fmt.Println("7. Confirm All Pending Orders")
	fmt.Println("8. Logout")

	choice := readInput("\nEnter your choice: ")
	executeEmployeeOperation(choice)
}

// Handle operation from Employee
func executeEmployeeOperation(choice string) bool {
	switch choice {
	case "0":
		fetchMenu()
	case "1":
		fmt.Println("Checking All Orders")
		fetchAllOrders()
		waitUntilPress()
	case "2":
		fmt.Println("Checking Order by ID")
		orderID := readInput("\nEnter Order ID: ")
		fetchOrderByID(orderID)
		waitUntilPress()
	case "3":
		fetchMenu()
		fmt.Println("\nCreate Order")
		payload := populateOrderRequest("")
		if payload != nil {
			createOrder(*payload)
			waitUntilPress()
		}
	case "4":
		fetchMenu()
		fmt.Println("\nUpdate Order")
		orderID := readInput("\nSearch Order by ID: ")
		fetchOrderByID(orderID)

		payload := populateOrderRequest(orderID)
		if payload != nil {
			updateOrder(*payload)
			waitUntilPress()
		}
	case "5":
		fmt.Println("Cancel Order")
		orderID := readInput("\nEnter order ID to Cancel: ")
		fetchOrderByID(orderID)
		triggerAction := readInput("\nAre you sure? y/n:")
		if triggerAction == "Y" || triggerAction == "y" {
			cancelOrder(orderID)
			waitUntilPress()
		}

	case "6":
		fmt.Println("Confirm Order")
		orderID := readInput("\nEnter order ID to Confirm: ")
		fetchOrderByID(orderID)
		triggerAction := readInput("\nAre you sure? y/n:")
		if triggerAction == "Y" || triggerAction == "y" {
			confirmOrder(orderID)
			waitUntilPress()
		}

	case "7":
		fmt.Println("Confirm All Pending Orders")
		triggerAction := readInput("\nAll Order will be send to the Kitchen, Are you sure? y/n:")
		if triggerAction == "Y" || triggerAction == "y" {
			confirmAllPreOrder()
			waitUntilPress()
		}
	case "8":
		fmt.Println("Log Out")
		logout()
		currentMenu = PublicMenu
		displayGuestMenu()

	default:
		fmt.Println("Invalid option, please try again.")
	}
	return true
}

// Function to read user Input
func readInput(str string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(str)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Function to wait the process until press ENTER
func waitUntilPress() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press ENTER to Continue...")
	reader.ReadString('\n')
	return
}
