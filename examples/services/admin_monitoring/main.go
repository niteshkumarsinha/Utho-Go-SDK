package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// 1. Account Info
	acc, err := client.Account.GetInfo()
	if err == nil {
		fmt.Println("--- Account Information ---")
		fmt.Printf("Email: %s\n", acc.Email)
		fmt.Printf("Name: %s %s\n", acc.FirstName, acc.LastName)
		fmt.Printf("Balance: %s %s\n", acc.Balance, acc.Currency)
	}

	// 2. Monitoring
	fmt.Println("\nListing Alert Policies...")
	policies, _ := client.Monitoring.ListAlertPolicies()
	for _, p := range policies {
		fmt.Printf("- Policy: %s (Type: %s) [Active: %v]\n", p.Label, p.Resource, p.Active)
	}
}
