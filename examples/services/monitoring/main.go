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

	fmt.Println("=== Monitoring Example ===\n")

	// 1. List all alert policies
	fmt.Println("1. Listing all alert policies...")
	policies, err := client.Monitoring.ListAlertPolicies()
	if err != nil {
		log.Fatalf("Error listing alert policies: %v", err)
	}
	fmt.Printf("Found %d alert policies\n", len(policies))
	for _, policy := range policies {
		fmt.Printf("  - ID: %s, Label: %s, Active: %t, Resource: %s\n",
			policy.ID, policy.Label, policy.Active, policy.Resource)
	}

	// 2. Create an alert policy (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating an alert policy...")
		createParams := monitoring.CreateAlertPolicyParams{
			Label:        "high-cpu-alert",
			ResourceType: "cloudserver",
			Contacts:     []string{"contact-id-1"},
		}
		createParams.Thresholds.CPU = 80
		createParams.Thresholds.RAM = 90

		err = client.Monitoring.CreateAlertPolicy(createParams)
		if err != nil {
			log.Printf("Error creating alert policy: %v", err)
		} else {
			fmt.Println("Alert policy created successfully")
		}
	*/

	// 3. Delete an alert policy (COMMENTED OUT)
	/*
		if len(policies) > 0 {
			policyID := policies[0].ID
			fmt.Printf("\n3. Deleting alert policy %s...\n", policyID)
			err = client.Monitoring.DeleteAlertPolicy(policyID)
			if err != nil {
				log.Printf("Error deleting alert policy: %v", err)
			} else {
				fmt.Println("Alert policy deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
