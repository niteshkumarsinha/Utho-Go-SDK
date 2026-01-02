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

	fmt.Println("=== Snapshots Example ===\n")

	// 1. List all snapshots
	fmt.Println("1. Listing all snapshots...")
	snaps, err := client.Snapshots.List()
	if err != nil {
		log.Fatalf("Error listing snapshots: %v", err)
	}
	fmt.Printf("Found %d snapshots\n", len(snaps))
	for _, snap := range snaps {
		fmt.Printf("  - ID: %s, Name: %s, Size: %s, Status: %s\n",
			snap.ID, snap.Name, snap.Size, snap.Status)
	}

	// 2. Create a snapshot (COMMENTED OUT)
	/*
		fmt.Println("\n2. Creating a snapshot...")
		createParams := snapshots.CreateParams{
			CloudID: "your-cloud-server-id",
			Name:    "my-snapshot",
		}
		err = client.Snapshots.Create(createParams)
		if err != nil {
			log.Printf("Error creating snapshot: %v", err)
		} else {
			fmt.Println("Snapshot creation initiated")
		}
	*/

	// 3. Delete a snapshot (COMMENTED OUT)
	/*
		if len(snaps) > 0 {
			cloudID := "your-cloud-server-id"
			snapshotID := snaps[0].ID
			fmt.Printf("\n3. Deleting snapshot %s...\n", snapshotID)
			err = client.Snapshots.Delete(cloudID, snapshotID)
			if err != nil {
				log.Printf("Error deleting snapshot: %v", err)
			} else {
				fmt.Println("Snapshot deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
