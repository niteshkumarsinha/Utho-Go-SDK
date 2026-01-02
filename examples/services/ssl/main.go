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

	fmt.Println("=== SSL Certificate Example ===\n")

	// 1. List all SSL certificates
	fmt.Println("1. Listing all SSL certificates...")
	certs, err := client.SSL.List()
	if err != nil {
		log.Fatalf("Error listing certificates: %v", err)
	}
	fmt.Printf("Found %d certificates\n", len(certs))
	for _, cert := range certs {
		fmt.Printf("  - ID: %s, Name: %s, Type: %s\n", cert.ID, cert.Name, cert.Type)
	}

	// 2. Upload an SSL certificate (COMMENTED OUT)
	/*
		fmt.Println("\n2. Uploading an SSL certificate...")
		createParams := ssl.CreateParams{
			Name:        "my-ssl-cert",
			Certificate: "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
			PrivateKey:  "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----",
			Chain:       "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
		}
		err = client.SSL.Create(createParams)
		if err != nil {
			log.Printf("Error uploading certificate: %v", err)
		} else {
			fmt.Println("Certificate uploaded successfully")
		}
	*/

	// 3. Delete an SSL certificate (COMMENTED OUT)
	/*
		if len(certs) > 0 {
			certID := certs[0].ID
			fmt.Printf("\n3. Deleting certificate %s...\n", certID)
			err = client.SSL.Delete(certID)
			if err != nil {
				log.Printf("Error deleting certificate: %v", err)
			} else {
				fmt.Println("Certificate deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
