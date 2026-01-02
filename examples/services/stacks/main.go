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

	fmt.Println("=== Stacks (Automation Templates) Example ===\n")

	// 1. List all stacks
	fmt.Println("1. Listing all stacks...")
	stacks, err := client.Stacks.List()
	if err != nil {
		log.Fatalf("Error listing stacks: %v", err)
	}
	fmt.Printf("Found %d stacks\n", len(stacks))
	for _, stack := range stacks {
		fmt.Printf("  - ID: %s, Title: %s, Status: %s\n", stack.ID, stack.Title, stack.Status)
	}

	// 2. Create a stack (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a stack...")
		createParams := stacks.CreateParams{
			Title:       "My Automation Stack",
			Description: "Automated server setup",
			Images:      "ubuntu-20.04-x64",
			Script:      "#!/bin/bash\napt-get update\napt-get install -y nginx",
			Status:      "1",
		}
		err = client.Stacks.Create(createParams)
		if err != nil {
			log.Printf("Error creating stack: %v", err)
		} else {
			fmt.Println("Stack created successfully")
		}
	*/

	// 3. Update a stack (COMMENTED OUT)
	/*
		if len(stacks) > 0 {
			stackID := stacks[0].ID
			fmt.Printf("\n3. Updating stack %s...\n", stackID)
			updateParams := stacks.UpdateParams{
				Title:       "Updated Stack Title",
				Description: "Updated description",
			}
			err = client.Stacks.Update(stackID, updateParams)
			if err != nil {
				log.Printf("Error updating stack: %v", err)
			} else {
				fmt.Println("Stack updated successfully")
			}
		}
	*/

	// 4. Delete a stack (COMMENTED OUT)
	/*
		if len(stacks) > 0 {
			stackID := stacks[0].ID
			fmt.Printf("\n4. Deleting stack %s...\n", stackID)
			err = client.Stacks.Delete(stackID)
			if err != nil {
				log.Printf("Error deleting stack: %v", err)
			} else {
				fmt.Println("Stack deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
