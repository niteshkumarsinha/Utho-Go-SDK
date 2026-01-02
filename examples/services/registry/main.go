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

	fmt.Println("=== Container Registry Example ===\n")

	// 1. List all registries
	fmt.Println("1. Listing all container registries...")
	registries, err := client.Registry.List()
	if err != nil {
		log.Fatalf("Error listing registries: %v", err)
	}
	fmt.Printf("Found %d registries\n", len(registries))
	for _, reg := range registries {
		fmt.Printf("  - ID: %s, Name: %s, URL: %s, Status: %s\n", reg.ID, reg.Name, reg.URL, reg.Status)
	}

	// 2. Create a registry (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a container registry...")
		createParams := registry.CreateParams{
			Name:   "my-registry",
			DCSlug: "inmumbaizone2",
		}
		err = client.Registry.Create(createParams)
		if err != nil {
			log.Printf("Error creating registry: %v", err)
		} else {
			fmt.Println("Registry created successfully")
		}
	*/

	// 3. Delete a registry (COMMENTED OUT)
	/*
		if len(registries) > 0 {
			regID := registries[0].ID
			fmt.Printf("\n3. Deleting registry %s...\n", regID)
			err = client.Registry.Delete(regID)
			if err != nil {
				log.Printf("Error deleting registry: %v", err)
			} else {
				fmt.Println("Registry deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
