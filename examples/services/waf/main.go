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

	fmt.Println("=== WAF (Web Application Firewall) Example ===\n")

	// 1. List all WAF instances
	fmt.Println("1. Listing all WAF instances...")
	wafs, err := client.WAF.List()
	if err != nil {
		log.Fatalf("Error listing WAFs: %v", err)
	}
	fmt.Printf("Found %d WAF instances\n", len(wafs))
	for _, w := range wafs {
		fmt.Printf("  - ID: %s, Name: %s, Status: %s\n", w.ID, w.Name, w.Status)
	}

	// 2. Create a WAF instance (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a WAF instance...")
		createParams := waf.CreateParams{
			Name:   "my-waf",
			DCSlug: "inmumbaizone2",
		}
		err = client.WAF.Create(createParams)
		if err != nil {
			log.Printf("Error creating WAF: %v", err)
		} else {
			fmt.Println("WAF created successfully")
		}
	*/

	// 3. Attach WAF to a resource (COMMENTED OUT)
	/*
		if len(wafs) > 0 {
			wafID := wafs[0].ID
			fmt.Printf("\n3. Attaching WAF %s to a load balancer...\n", wafID)
			attachParams := waf.AttachParams{
				ResourceID:   "your-loadbalancer-id",
				ResourceType: "loadbalancer",
			}
			err = client.WAF.Attach(wafID, attachParams)
			if err != nil {
				log.Printf("Error attaching WAF: %v", err)
			} else {
				fmt.Println("WAF attached successfully")
			}
		}
	*/

	// 4. Detach WAF (COMMENTED OUT)
	/*
		if len(wafs) > 0 {
			wafID := wafs[0].ID
			fmt.Printf("\n4. Detaching WAF %s...\n", wafID)
			err = client.WAF.Detach(wafID)
			if err != nil {
				log.Printf("Error detaching WAF: %v", err)
			} else {
				fmt.Println("WAF detached successfully")
			}
		}
	*/

	// 5. Delete a WAF instance (COMMENTED OUT)
	/*
		if len(wafs) > 0 {
			wafID := wafs[0].ID
			fmt.Printf("\n5. Deleting WAF %s...\n", wafID)
			err = client.WAF.Delete(wafID)
			if err != nil {
				log.Printf("Error deleting WAF: %v", err)
			} else {
				fmt.Println("WAF deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
