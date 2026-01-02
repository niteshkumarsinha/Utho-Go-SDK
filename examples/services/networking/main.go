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

	fmt.Println("=== Networking (DNS & Firewalls) Example ===\n")

	// 1. List DNS domains
	fmt.Println("1. Listing DNS domains...")
	domains, err := client.Networking.ListDomains()
	if err != nil {
		log.Fatalf("Error listing domains: %v", err)
	}
	fmt.Printf("Found %d domains\n", len(domains))
	for _, domain := range domains {
		fmt.Printf("  - ID: %s, Name: %s\n", domain.ID, domain.Name)
	}

	// 2. List firewalls
	fmt.Println("\n2. Listing firewalls...")
	firewalls, err := client.Networking.ListFirewalls()
	if err != nil {
		log.Fatalf("Error listing firewalls: %v", err)
	}
	fmt.Printf("Found %d firewalls\n", len(firewalls))
	for _, fw := range firewalls {
		fmt.Printf("  - ID: %s, Name: %s, Status: %s\n", fw.ID, fw.Name, fw.Status)
	}

	// 3. Create a DNS domain (COMMENTED OUT)
	/*
		fmt.Println("\n3. Creating a DNS domain...")
		createDomainParams := networking.CreateDomainParams{
			Domain: "example.com",
		}
		err = client.Networking.CreateDomain(createDomainParams)
		if err != nil {
			log.Printf("Error creating domain: %v", err)
		} else {
			fmt.Println("Domain created successfully")
		}
	*/

	// 4. Create a firewall (COMMENTED OUT)
	/*
		fmt.Println("\n4. Creating a firewall...")
		createFWParams := networking.CreateFirewallParams{
			Name: "my-firewall",
		}
		err = client.Networking.CreateFirewall(createFWParams)
		if err != nil {
			log.Printf("Error creating firewall: %v", err)
		} else {
			fmt.Println("Firewall created successfully")
		}
	*/

	// 5. Delete a domain (COMMENTED OUT)
	/*
		if len(domains) > 0 {
			domainName := domains[0].Name
			fmt.Printf("\n5. Deleting domain %s...\n", domainName)
			err = client.Networking.DeleteDomain(domainName)
			if err != nil {
				log.Printf("Error deleting domain: %v", err)
			} else {
				fmt.Println("Domain deleted successfully")
			}
		}
	*/

	// 6. Delete a firewall (COMMENTED OUT)
	/*
		if len(firewalls) > 0 {
			fwID := firewalls[0].ID
			fmt.Printf("\n6. Deleting firewall %s...\n", fwID)
			err = client.Networking.DeleteFirewall(fwID)
			if err != nil {
				log.Printf("Error deleting firewall: %v", err)
			} else {
				fmt.Println("Firewall deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
