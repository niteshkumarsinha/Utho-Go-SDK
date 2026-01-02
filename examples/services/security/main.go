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

	fmt.Println("=== Security (SSH & API Keys) Example ===\n")

	// 1. List SSH keys
	fmt.Println("1. Listing SSH keys...")
	sshKeys, err := client.Security.ListSSHKeys()
	if err != nil {
		log.Fatalf("Error listing SSH keys: %v", err)
	}
	fmt.Printf("Found %d SSH keys\n", len(sshKeys))
	for _, key := range sshKeys {
		fmt.Printf("  - ID: %s, Name: %s\n", key.ID, key.Name)
	}

	// 2. List API keys
	fmt.Println("\n2. Listing API keys...")
	apiKeys, err := client.Security.ListAPIKeys()
	if err != nil {
		log.Fatalf("Error listing API keys: %v", err)
	}
	fmt.Printf("Found %d API keys\n", len(apiKeys))
	for _, key := range apiKeys {
		fmt.Printf("  - ID: %s, Label: %s\n", key.ID, key.Label)
	}

	// 3. Import SSH key (COMMENTED OUT)
	/*
		fmt.Println("\n3. Importing SSH key...")
		importParams := security.ImportSSHKeyParams{
			Name: "my-ssh-key",
			Key:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC...",
		}
		err = client.Security.ImportSSHKey(importParams)
		if err != nil {
			log.Printf("Error importing SSH key: %v", err)
		} else {
			fmt.Println("SSH key imported successfully")
		}
	*/

	// 4. Generate API key (COMMENTED OUT)
	/*
		fmt.Println("\n4. Generating API key...")
		genParams := security.GenerateAPIKeyParams{
			Label: "my-api-key",
		}
		apiKey, err := client.Security.GenerateAPIKey(genParams)
		if err != nil {
			log.Printf("Error generating API key: %v", err)
		} else {
			fmt.Printf("API key generated: %s\n", apiKey.Key)
		}
	*/

	// 5. Delete SSH key (COMMENTED OUT)
	/*
		if len(sshKeys) > 0 {
			keyID := sshKeys[0].ID
			fmt.Printf("\n5. Deleting SSH key %s...\n", keyID)
			err = client.Security.DeleteSSHKey(keyID)
			if err != nil {
				log.Printf("Error deleting SSH key: %v", err)
			} else {
				fmt.Println("SSH key deleted successfully")
			}
		}
	*/

	// 6. Delete API key (COMMENTED OUT)
	/*
		if len(apiKeys) > 0 {
			keyID := apiKeys[0].ID
			fmt.Printf("\n6. Deleting API key %s...\n", keyID)
			err = client.Security.DeleteAPIKey(keyID)
			if err != nil {
				log.Printf("Error deleting API key: %v", err)
			} else {
				fmt.Println("API key deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
