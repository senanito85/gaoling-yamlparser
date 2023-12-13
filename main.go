package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Company represents the company information.
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Customer represents the customer information.
type Customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Number  string  `json:"number"`
	Email   string  `json:"email"`
	Company Company `json:"company"`
}

// CustomersWrapper represents the top-level structure of the JSON file.
type CustomersWrapper struct {
	Customers []Customer `json:"customers"`
}

func main() {
	// Check if a JSON file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go dataset.json")
		os.Exit(1)
	}

	// Get the JSON file path from command-line arguments
	jsonFilePath := os.Args[1]

	// Read the content of the JSON file
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		os.Exit(1)
	}

	// Parse the JSON data into a struct with a Customers field
	var customersWrapper CustomersWrapper
	err = json.Unmarshal(jsonData, &customersWrapper)
	if err != nil {
		fmt.Printf("Error parsing JSON data: %v\n", err)
		os.Exit(1)
	}

	// Print each customer and company information
	for _, customer := range customersWrapper.Customers {
		fmt.Printf("Customer ID: %d\n", customer.ID)
		fmt.Printf("  Name: %s\n", customer.Name)
		fmt.Printf("  Number: %s\n", customer.Number)
		fmt.Printf("  Email: %s\n", customer.Email)

		fmt.Printf("  Company ID: %d\n", customer.Company.ID)
		fmt.Printf("    Name: %s\n", customer.Company.Name)
		fmt.Printf("    Address: %s\n", customer.Company.Address)

		fmt.Println("--------------------")
	}
}
