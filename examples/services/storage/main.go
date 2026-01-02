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

	fmt.Println("=== Storage (EBS) Example ===\n")

	// 1. List all EBS volumes
	fmt.Println("1. Listing all EBS volumes...")
	volumes, err := client.Storage.List()
	if err != nil {
		log.Fatalf("Error listing volumes: %v", err)
	}
	fmt.Printf("Found %d volumes\n", len(volumes))
	for _, vol := range volumes {
		fmt.Printf("  - ID: %s, Name: %s, Size: %s, Status: %s, Attached: %s\n",
			vol.ID, vol.Name, vol.Size, vol.Status, vol.Attached)
	}

	// 2. Create a new EBS volume (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a new EBS volume...")
		createParams := storage.CreateParams{
			Name:     "my-ebs-volume",
			DCSlug:   "inmumbaizone2",
			Size:     "50", // GB
			DiskType: "ssd",
		}
		createResp, err := client.Storage.Create(createParams)
		if err != nil {
			log.Printf("Error creating volume: %v", err)
		} else {
			fmt.Printf("Volume created. EBS ID: %s\n", createResp.EBSID)
		}
	*/

	// 3. Attach volume to a server (COMMENTED OUT)
	/*
		if len(volumes) > 0 {
			ebsID := volumes[0].ID
			fmt.Printf("\n3. Attaching volume %s to server...\n", ebsID)
			attachParams := storage.AttachParams{
				ServerID: "your-server-id-here",
			}
			err = client.Storage.Attach(ebsID, attachParams)
			if err != nil {
				log.Printf("Error attaching volume: %v", err)
			} else {
				fmt.Println("Volume attached successfully")
			}
		}
	*/

	// 4. Detach volume (COMMENTED OUT)
	/*
		if len(volumes) > 0 {
			ebsID := volumes[0].ID
			fmt.Printf("\n4. Detaching volume %s...\n", ebsID)
			err = client.Storage.Detach(ebsID)
			if err != nil {
				log.Printf("Error detaching volume: %v", err)
			} else {
				fmt.Println("Volume detached successfully")
			}
		}
	*/

	// 5. Delete volume (COMMENTED OUT)
	/*
		if len(volumes) > 0 {
			ebsID := volumes[0].ID
			fmt.Printf("\n5. Deleting volume %s...\n", ebsID)
			err = client.Storage.Delete(ebsID)
			if err != nil {
				log.Printf("Error deleting volume: %v", err)
			} else {
				fmt.Println("Volume deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
