package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/stacks"
	"github.com/niteshkumarsinha/utho-sdk-go/services/transfer"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// 1. Stacks
	fmt.Println("Creating a new automation stack...")
	err := client.Stacks.Create(stacks.CreateParams{
		Title:       "My Web Stack",
		Description: "Auto-deploy Ubuntu with Nginx",
		Images:      "ubuntu-22.04-x86_64",
		Script:      "#!/bin/bash\napt update && apt install -y nginx",
	})
	if err != nil {
		fmt.Printf("Error creating stack: %v\n", err)
	}

	fmt.Println("\nListing stacks...")
	sList, _ := client.Stacks.List()
	for _, s := range sList {
		fmt.Printf("- %s (ID: %s)\n", s.Title, s.ID)
	}

	// 2. Transfer
	fmt.Println("\nInitiating resource transfer...")
	err = client.Transfer.Initiate("cloud", "12345")
	if err != nil {
		fmt.Printf("Error initiating transfer: %v\n", err)
	}

	fmt.Println("\nProcessing received resource...")
	err = client.Transfer.Receive(transfer.ReceiveParams{
		ID:    "123456",
		Token: "abc-def-ghi",
		Type:  "cloud",
	})
	if err != nil {
		fmt.Printf("Error receiving transfer: %v\n", err)
	}
}
