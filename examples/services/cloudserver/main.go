package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("UTHO_API_KEY")
	if apiKey == "" {
		log.Fatal("UTHO_API_KEY environment variable is required")
	}

	// Create client
	client, err := utho.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	fmt.Println("=== Cloud Server Example ===\n")

	// 1. List all cloud servers
	fmt.Println("1. Listing all cloud servers...")
	servers, err := client.CloudServer.List()
	if err != nil {
		log.Fatalf("Error listing servers: %v", err)
	}
	fmt.Printf("Found %d servers\n", len(servers))
	for _, server := range servers {
		fmt.Printf("  - ID: %s, Hostname: %s, IP: %s, Status: %s\n",
			server.ID, server.Hostname, server.IP, server.Status)
	}

	// 2. Get details of a specific server (if any exist)
	if len(servers) > 0 {
		fmt.Println("\n2. Getting details of first server...")
		serverID := servers[0].ID
		server, err := client.CloudServer.Get(serverID)
		if err != nil {
			log.Printf("Error getting server: %v", err)
		} else {
			fmt.Printf("Server Details:\n")
			fmt.Printf("  ID: %s\n", server.ID)
			fmt.Printf("  Hostname: %s\n", server.Hostname)
			fmt.Printf("  IP: %s\n", server.IP)
			fmt.Printf("  Status: %s\n", server.Status)
			fmt.Printf("  DC: %s\n", server.DC)
		}
	}

	// 3. Deploy a new cloud server (COMMENTED OUT - uncomment to create)
	/*
		fmt.Println("\n3. Deploying a new cloud server...")
		deployParams := cloudserver.DeployParams{
			DCSlug:         "inmumbaizone2",
			PlanID:         "10045",
			BillingCycle:   "hourly",
			Auth:           "password",
			EnablePublicIP: "1",
			Image:          "ubuntu-20.04-x64",
			Cloud: []cloudserver.InstanceConfig{
				{Hostname: "test-server-1"},
			},
		}
		deployResp, err := client.CloudServer.Deploy(deployParams)
		if err != nil {
			log.Printf("Error deploying server: %v", err)
		} else {
			fmt.Printf("Server deployment initiated. Order ID: %s\n", deployResp.OrderID)
		}
	*/

	// 4. Power operations (COMMENTED OUT - uncomment to use)
	/*
		if len(servers) > 0 {
			serverID := servers[0].ID

			// Power off
			fmt.Println("\n4. Powering off server...")
			err = client.CloudServer.PowerOff(serverID)
			if err != nil {
				log.Printf("Error powering off: %v", err)
			}

			// Power on
			fmt.Println("Powering on server...")
			err = client.CloudServer.PowerOn(serverID)
			if err != nil {
				log.Printf("Error powering on: %v", err)
			}

			// Hard reboot
			fmt.Println("Hard rebooting server...")
			err = client.CloudServer.HardReboot(serverID)
			if err != nil {
				log.Printf("Error rebooting: %v", err)
			}
		}
	*/

	// 5. Delete a server (COMMENTED OUT - uncomment to delete)
	/*
		if len(servers) > 0 {
			serverID := servers[0].ID
			fmt.Printf("\n5. Deleting server %s...\n", serverID)
			err = client.CloudServer.Delete(serverID)
			if err != nil {
				log.Printf("Error deleting server: %v", err)
			} else {
				fmt.Println("Server deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
