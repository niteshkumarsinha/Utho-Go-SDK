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

	fmt.Println("=== ISO Example ===\n")

	// 1. List all custom ISOs
	fmt.Println("1. Listing all custom ISOs...")
	isos, err := client.ISO.List()
	if err != nil {
		log.Fatalf("Error listing ISOs: %v", err)
	}
	fmt.Printf("Found %d ISOs\n", len(isos))
	for _, iso := range isos {
		fmt.Printf("  - ID: %s, Name: %s, Status: %s\n",
			iso.ID, iso.Name, iso.Status)
	}

	// 2. Create a custom ISO (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a custom ISO...")
		createParams := iso.CreateParams{
			Name:   "my-custom-iso",
			URL:    "https://example.com/path/to/iso",
			DCSlug: "inmumbaizone2",
		}
		err = client.ISO.Create(createParams)
		if err != nil {
			log.Printf("Error creating ISO: %v", err)
		} else {
			fmt.Println("ISO creation initiated")
		}
	*/

	// 3. Delete an ISO (COMMENTED OUT)
	/*
		if len(isos) > 0 {
			isoID := isos[0].ID
			fmt.Printf("\n3. Deleting ISO %s...\n", isoID)
			err = client.ISO.Delete(isoID)
			if err != nil {
				log.Printf("Error deleting ISO: %v", err)
			} else {
				fmt.Println("ISO deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
