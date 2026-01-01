package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// 1. SSH Keys
	fmt.Println("Listing SSH keys...")
	keys, _ := client.Security.ListSSHKeys()
	for _, k := range keys {
		fmt.Printf("- Key: %s (ID: %s)\n", k.Name, k.ID)
	}

	// 2. SSL Certificates
	fmt.Println("\nListing SSL certificates...")
	certs, _ := client.SSL.List()
	for _, c := range certs {
		fmt.Printf("- Cert: %s (Type: %s)\n", c.Name, c.Type)
	}

	// 3. WAF
	fmt.Println("\nListing WAF instances...")
	wafs, _ := client.WAF.List()
	for _, w := range wafs {
		fmt.Printf("- WAF: %s [%s]\n", w.Name, w.Status)
	}
}
