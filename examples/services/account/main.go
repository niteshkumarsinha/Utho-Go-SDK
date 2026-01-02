package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	if apiKey == "" {
		log.Fatal("UTHO_API_KEY environment variable is required")
	}

	client, err := utho.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	fmt.Println("=== Account Example ===\n")

	// Get account information
	fmt.Println("Getting account information...")
	info, err := client.Account.GetInfo()
	if err != nil {
		log.Fatalf("Error getting account info: %v", err)
	}

	fmt.Printf("Account Details:\n")
	fmt.Printf("  Email: %s\n", info.Email)
	fmt.Printf("  Name: %s %s\n", info.FirstName, info.LastName)
	fmt.Printf("  Status: %s\n", info.Status)
	fmt.Printf("  Currency: %s\n", info.Currency)
	fmt.Printf("  Balance: %s\n", info.Balance)

	fmt.Println("\n=== Example completed ===")
}
