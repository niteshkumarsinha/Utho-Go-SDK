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

	fmt.Println("=== Load Balancer Example ===\n")

	// 1. List all load balancers
	fmt.Println("1. Listing all load balancers...")
	lbs, err := client.LoadBalancer.List()
	if err != nil {
		log.Fatalf("Error listing load balancers: %v", err)
	}
	fmt.Printf("Found %d load balancers\n", len(lbs))
	for _, lb := range lbs {
		fmt.Printf("  - ID: %s, Name: %s, Type: %s, IP: %s\n", lb.ID, lb.Name, lb.Type, lb.IP)
	}

	// 2. Create a load balancer (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a load balancer...")
		createParams := loadbalancer.CreateParams{
			DCSlug: "inmumbaizone2",
			Name:   "my-lb",
			Type:   "http",
		}
		createResp, err := client.LoadBalancer.Create(createParams)
		if err != nil {
			log.Printf("Error creating load balancer: %v", err)
		} else {
			fmt.Printf("Load balancer created. ID: %s\n", createResp.LBID)
		}
	*/

	// 3. Update a load balancer (COMMENTED OUT)
	/*
		if len(lbs) > 0 {
			lbID := lbs[0].ID
			fmt.Printf("\n3. Updating load balancer %s...\n", lbID)
			updateParams := loadbalancer.UpdateParams{
				Name: "updated-lb-name",
			}
			err = client.LoadBalancer.Update(lbID, updateParams)
			if err != nil {
				log.Printf("Error updating load balancer: %v", err)
			} else {
				fmt.Println("Load balancer updated successfully")
			}
		}
	*/

	// 4. Delete a load balancer (COMMENTED OUT)
	/*
		if len(lbs) > 0 {
			lbID := lbs[0].ID
			fmt.Printf("\n4. Deleting load balancer %s...\n", lbID)
			err = client.LoadBalancer.Delete(lbID)
			if err != nil {
				log.Printf("Error deleting load balancer: %v", err)
			} else {
				fmt.Println("Load balancer deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
