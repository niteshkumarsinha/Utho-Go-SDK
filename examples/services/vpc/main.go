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

	fmt.Println("=== VPC Example ===\n")

	// 1. List all VPCs
	fmt.Println("1. Listing all VPCs...")
	vpcs, err := client.VPC.List()
	if err != nil {
		log.Fatalf("Error listing VPCs: %v", err)
	}
	fmt.Printf("Found %d VPCs\n", len(vpcs))
	for _, v := range vpcs {
		fmt.Printf("  - ID: %s, Name: %s, Network: %s\n", v.ID, v.Name, v.Network)
	}

	// 2. Create a VPC (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a VPC...")
		createParams := vpc.CreateParams{
			Name:   "my-vpc",
			Region: "inmumbaizone2",
			Range:  "10.0.0.0/16",
		}
		err = client.VPC.Create(createParams)
		if err != nil {
			log.Printf("Error creating VPC: %v", err)
		} else {
			fmt.Println("VPC created successfully")
		}
	*/

	// 3. Delete a VPC (COMMENTED OUT)
	/*
		if len(vpcs) > 0 {
			vpcID := vpcs[0].ID
			fmt.Printf("\n3. Deleting VPC %s...\n", vpcID)
			err = client.VPC.Delete(vpcID)
			if err != nil {
				log.Printf("Error deleting VPC: %v", err)
			} else {
				fmt.Println("VPC deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
