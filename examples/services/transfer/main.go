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

	fmt.Println("=== Transfer Example ===\n")

	// 1. Initiate a resource transfer (COMMENTED OUT)
	/*
		fmt.Println("1. Initiating a resource transfer...")
		resourceType := "cloudserver"
		resourceID := "your-server-id"

		err = client.Transfer.Initiate(resourceType, resourceID)
		if err != nil {
			log.Printf("Error initiating transfer: %v", err)
		} else {
			fmt.Println("Transfer initiated successfully")
			fmt.Println("A transfer token will be sent to the recipient")
		}
	*/

	// 2. Receive a transferred resource (COMMENTED OUT)
	/*
		fmt.Println("\n2. Receiving a transferred resource...")
		receiveParams := transfer.ReceiveParams{
			ID:    "resource-id",
			Token: "transfer-token-from-sender",
			Type:  "cloudserver",
		}

		err = client.Transfer.Receive(receiveParams)
		if err != nil {
			log.Printf("Error receiving transfer: %v", err)
		} else {
			fmt.Println("Resource transfer received successfully")
		}
	*/

	fmt.Println("\n=== Example completed ===")
	fmt.Println("\nNote: Transfer operations require coordination between two accounts.")
	fmt.Println("The sender initiates the transfer and provides a token to the recipient.")
}
