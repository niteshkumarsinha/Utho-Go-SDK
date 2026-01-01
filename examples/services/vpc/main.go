package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/vpc"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	fmt.Println("Creating a new VPC...")

	params := vpc.CreateParams{
		Name:      "my-prod-vpc",
		Region:    "inmumbaizone2",
		IPv4Range: "10.0.0.0/16",
	}

	resp, err := client.VPC.Create(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("VPC Created: %s\n", resp.VPCID)
	}

	fmt.Println("\nListing VPCs...")
	vpcs, _ := client.VPC.List()
	for _, v := range vpcs {
		fmt.Printf("- %s: %s (%s)\n", v.Name, v.IPv4Range, v.Region)
	}
}
