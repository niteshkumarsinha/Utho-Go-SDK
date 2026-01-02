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

	fmt.Println("=== Autoscaling Example ===\n")

	// 1. List all autoscaling groups
	fmt.Println("1. Listing all autoscaling groups...")
	groups, err := client.Autoscaling.List()
	if err != nil {
		log.Fatalf("Error listing autoscaling groups: %v", err)
	}
	fmt.Printf("Found %d autoscaling groups\n", len(groups))
	for _, group := range groups {
		fmt.Printf("  - ID: %s, Name: %s, Min: %d, Max: %d, Status: %s\n",
			group.ID, group.Name, group.MinSize, group.MaxSize, group.Status)
	}

	// 2. Create an autoscaling group (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating an autoscaling group...")
		createParams := autoscaling.CreateParams{
			Name:    "my-as-group",
			MinSize: 1,
			MaxSize: 5,
			Image:   "ubuntu-20.04-x64",
			Plan:    "10045",
			Script:  "#!/bin/bash\necho 'Hello World'",
		}
		err = client.Autoscaling.Create(createParams)
		if err != nil {
			log.Printf("Error creating autoscaling group: %v", err)
		} else {
			fmt.Println("Autoscaling group created successfully")
		}
	*/

	// 3. Update an autoscaling group (COMMENTED OUT)
	/*
		if len(groups) > 0 {
			groupID := groups[0].ID
			fmt.Printf("\n3. Updating autoscaling group %s...\n", groupID)
			updateParams := autoscaling.UpdateParams{
				MinSize: 2,
				MaxSize: 10,
			}
			err = client.Autoscaling.Update(groupID, updateParams)
			if err != nil {
				log.Printf("Error updating autoscaling group: %v", err)
			} else {
				fmt.Println("Autoscaling group updated successfully")
			}
		}
	*/

	// 4. Delete an autoscaling group (COMMENTED OUT)
	/*
		if len(groups) > 0 {
			groupID := groups[0].ID
			fmt.Printf("\n4. Deleting autoscaling group %s...\n", groupID)
			err = client.Autoscaling.Delete(groupID)
			if err != nil {
				log.Printf("Error deleting autoscaling group: %v", err)
			} else {
				fmt.Println("Autoscaling group deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
