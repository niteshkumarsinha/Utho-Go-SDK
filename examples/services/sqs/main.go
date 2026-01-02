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

	fmt.Println("=== SQS (Simple Queue Service) Example ===\n")

	// 1. List all SQS instances
	fmt.Println("1. Listing all SQS instances...")
	queues, err := client.SQS.List()
	if err != nil {
		log.Fatalf("Error listing SQS instances: %v", err)
	}
	fmt.Printf("Found %d SQS instances\n", len(queues))
	for _, queue := range queues {
		fmt.Printf("  - ID: %s, Name: %s, Status: %s\n", queue.ID, queue.Name, queue.Status)
	}

	// 2. Create an SQS instance (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating an SQS instance...")
		createParams := sqs.CreateParams{
			Name:   "my-queue",
			DCSlug: "inmumbaizone2",
		}
		err = client.SQS.Create(createParams)
		if err != nil {
			log.Printf("Error creating SQS instance: %v", err)
		} else {
			fmt.Println("SQS instance created successfully")
		}
	*/

	// 3. Delete an SQS instance (COMMENTED OUT)
	/*
		if len(queues) > 0 {
			queueID := queues[0].ID
			fmt.Printf("\n3. Deleting SQS instance %s...\n", queueID)
			err = client.SQS.Delete(queueID)
			if err != nil {
				log.Printf("Error deleting SQS instance: %v", err)
			} else {
				fmt.Println("SQS instance deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
