package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/cloudserver"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	fmt.Println("Deploying a new cloud server...")

	params := cloudserver.DeployParams{
		DCSlug:       "inmumbaizone2",
		PlanID:       "10045",
		BillingCycle: "hourly",
		Image:        "ubuntu-22.04-x86_64",
		Cloud: []cloudserver.InstanceConfig{
			{Hostname: "example-server"},
		},
	}

	resp, err := client.CloudServer.Deploy(params)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Success! Order ID: %s\n", resp.OrderID)

	fmt.Println("Listing active servers...")
	servers, _ := client.CloudServer.List()
	for _, s := range servers {
		fmt.Printf("- %s (%s) [%s]\n", s.Hostname, s.IP, s.Status)
	}
}
