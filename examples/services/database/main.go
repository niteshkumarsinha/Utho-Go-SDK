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

	fmt.Println("=== Database Example ===\n")

	// 1. List all database clusters
	fmt.Println("1. Listing all database clusters...")
	dbs, err := client.Database.List()
	if err != nil {
		log.Fatalf("Error listing databases: %v", err)
	}
	fmt.Printf("Found %d database clusters\n", len(dbs))
	for _, db := range dbs {
		fmt.Printf("  - ID: %s, Name: %s, Engine: %s, Status: %s\n", db.ID, db.Name, db.Engine, db.Status)
	}

	// 2. Get details of a specific database
	if len(dbs) > 0 {
		fmt.Println("\n2. Getting details of first database...")
		dbID := dbs[0].ID
		db, err := client.Database.Get(dbID)
		if err != nil {
			log.Printf("Error getting database: %v", err)
		} else {
			fmt.Printf("Database Details:\n")
			fmt.Printf("  ID: %s\n", db.ID)
			fmt.Printf("  Name: %s\n", db.Name)
			fmt.Printf("  Engine: %s\n", db.Engine)
			fmt.Printf("  Status: %s\n", db.Status)
		}
	}

	// 3. Create a database cluster (COMMENTED OUT)
	/*
		fmt.Println("\n3. Creating a database cluster...")
		createParams := database.CreateParams{
			Name:     "my-db-cluster",
			Engine:   "mysql",
			Version:  "8.0",
			DCSlug:   "inmumbaizone2",
			PlanID:   "10045",
			Replicas: 1,
		}
		err = client.Database.Create(createParams)
		if err != nil {
			log.Printf("Error creating database: %v", err)
		} else {
			fmt.Println("Database cluster creation initiated")
		}
	*/

	// 4. Delete a database cluster (COMMENTED OUT)
	/*
		if len(dbs) > 0 {
			dbID := dbs[0].ID
			fmt.Printf("\n4. Deleting database %s...\n", dbID)
			err = client.Database.Delete(dbID)
			if err != nil {
				log.Printf("Error deleting database: %v", err)
			} else {
				fmt.Println("Database deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
