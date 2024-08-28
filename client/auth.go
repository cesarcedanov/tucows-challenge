package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtToken, employeeUsername string

func login() {
	username := readInput("Username: ")
	password := readInput("Password: ")
	jsonData, err := json.Marshal(UserCredentials{username, password})
	if err != nil {
		log.Println("Error marshalling data:", err)
		return
	}

	resp, err := client.Post(url+"login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending data:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed.")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	tokenResponse := struct {
		Token string `json:"token"`
	}{}
	jsonErr := json.Unmarshal(body, &tokenResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	jwtToken = tokenResponse.Token
	employeeUsername = username
	fmt.Println("Successfully logged in.")
}

func logout() {
	jwtToken = ""
	employeeUsername = ""
}
