package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// DNS Management
	fmt.Println("Listing DNS domains...")
	domains, err := client.Networking.ListDomains()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, d := range domains {
			fmt.Printf("- %s (ID: %s)\n", d.Name, d.ID)
		}
	}

	// Firewall Management
	fmt.Println("\nListing Firewalls...")
	fws, err := client.Networking.ListFirewalls()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, fw := range fws {
			fmt.Printf("- %s [%s]\n", fw.Name, fw.Status)
		}
	}
}
