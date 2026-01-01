package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	// 1. Snapshots
	fmt.Println("Listing snapshots...")
	snaps, _ := client.Snapshots.List()
	for _, s := range snaps {
		fmt.Printf("- Snapshot: %s (of server %s) [%s]\n", s.Name, s.Hostname, s.Status)
	}

	// 2. Backups
	fmt.Println("\nListing backups...")
	backups, _ := client.Backups.List()
	for _, b := range backups {
		fmt.Printf("- Backup: %s (of server %s)\n", b.CreatedAt, b.Hostname)
	}

	// 3. Autoscaling
	fmt.Println("\nListing autoscaling groups...")
	groups, _ := client.Autoscaling.List()
	for _, g := range groups {
		fmt.Printf("- Group: %s (Min: %d, Max: %d) [%s]\n", g.Name, g.MinSize, g.MaxSize, g.Status)
	}

	// 4. ISO
	fmt.Println("\nListing custom ISOs...")
	isos, _ := client.ISO.List()
	for _, iso := range isos {
		fmt.Printf("- ISO: %s (Size: %s)\n", iso.Name, iso.Size)
	}
}
