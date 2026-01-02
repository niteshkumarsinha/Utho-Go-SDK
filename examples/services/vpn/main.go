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

	fmt.Println("=== VPN Example ===\n")

	// 1. List all VPN instances
	fmt.Println("1. Listing all VPN instances...")
	vpns, err := client.VPN.List()
	if err != nil {
		log.Fatalf("Error listing VPNs: %v", err)
	}
	fmt.Printf("Found %d VPN instances\n", len(vpns))
	for _, v := range vpns {
		fmt.Printf("  - ID: %s, Name: %s, IP: %s, Status: %s\n", v.ID, v.Name, v.IP, v.Status)
	}

	// 2. Create a VPN instance (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a VPN instance...")
		createParams := vpn.CreateParams{
			Name:   "my-vpn",
			DCSlug: "inmumbaizone2",
			Plan:   "10045",
		}
		err = client.VPN.Create(createParams)
		if err != nil {
			log.Printf("Error creating VPN: %v", err)
		} else {
			fmt.Println("VPN created successfully")
		}
	*/

	// 3. Delete a VPN instance (COMMENTED OUT)
	/*
		if len(vpns) > 0 {
			vpnID := vpns[0].ID
			fmt.Printf("\n3. Deleting VPN %s...\n", vpnID)
			err = client.VPN.Delete(vpnID)
			if err != nil {
				log.Printf("Error deleting VPN: %v", err)
			} else {
				fmt.Println("VPN deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
