package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// 1. Container Registry
	fmt.Println("Listing container registries...")
	regs, _ := client.Registry.List()
	for _, r := range regs {
		fmt.Printf("- Registry: %s (URL: %s) [%s]\n", r.Name, r.URL, r.Status)
	}

	// 2. SQS
	fmt.Println("\nListing SQS instances...")
	queues, _ := client.SQS.List()
	for _, q := range queues {
		fmt.Printf("- SQS: %s [%s]\n", q.Name, q.Status)
	}
}
