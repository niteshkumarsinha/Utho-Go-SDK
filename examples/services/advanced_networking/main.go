package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// 1. Load Balancers
	fmt.Println("Listing Load Balancers...")
	lbs, _ := client.LoadBalancer.List()
	for _, lb := range lbs {
		fmt.Printf("- LB: %s (IP: %s) [%s]\n", lb.Name, lb.IP, lb.Status)
	}

	// 2. VPN
	fmt.Println("\nListing VPN instances...")
	vpns, _ := client.VPN.List()
	for _, v := range vpns {
		fmt.Printf("- VPN: %s (IP: %s) [%s]\n", v.Name, v.IP, v.Status)
	}

	// 3. Storage (EBS)
	fmt.Println("\nListing EBS volumes...")
	volumes, _ := client.Storage.List()
	for _, vol := range volumes {
		fmt.Printf("- Volume: %s (%s GB) [Attached: %s]\n", vol.Name, vol.Size, vol.Attached)
	}
}
