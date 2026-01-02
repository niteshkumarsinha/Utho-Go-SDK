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

	fmt.Println("=== Backups Example ===\n")

	// 1. List all backups
	fmt.Println("1. Listing all backups...")
	backups, err := client.Backups.List()
	if err != nil {
		log.Fatalf("Error listing backups: %v", err)
	}
	fmt.Printf("Found %d backups\n", len(backups))
	for _, backup := range backups {
		fmt.Printf("  - ID: %s, Name: %s, Size: %s, Status: %s\n",
			backup.ID, backup.Name, backup.Size, backup.Status)
	}

	// 2. Delete a backup (COMMENTED OUT)
	/*
		if len(backups) > 0 {
			backupID := backups[0].ID
			fmt.Printf("\n2. Deleting backup %s...\n", backupID)
			err = client.Backups.Delete(backupID)
			if err != nil {
				log.Printf("Error deleting backup: %v", err)
			} else {
				fmt.Println("Backup deleted successfully")
			}
		}
	*/

	// 3. Restore a backup (COMMENTED OUT)
	/*
		if len(backups) > 0 {
			backupID := backups[0].ID
			cloudID := "your-cloud-server-id"
			fmt.Printf("\n3. Restoring backup %s to server %s...\n", backupID, cloudID)
			err = client.Backups.Restore(backupID, cloudID)
			if err != nil {
				log.Printf("Error restoring backup: %v", err)
			} else {
				fmt.Println("Backup restore initiated")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
