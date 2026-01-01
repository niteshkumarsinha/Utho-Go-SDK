package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/database"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	fmt.Println("Deploying a Managed MySQL Cluster...")

	params := database.CreateParams{
		DCSlug:         "inmumbaizone2",
		ClusterLabel:   "my-db-cluster",
		ClusterEngine:  "mysql",
		ClusterVersion: "8.0",
		Size:           "10045",
		ReplicaCount:   1,
	}

	resp, err := client.Database.Create(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("DB Cluster created: %s\n", resp.ClusterID)

	dbs, _ := client.Database.List()
	for _, db := range dbs {
		fmt.Printf("- %s (%s) [%s]\n", db.ClusterLabel, db.ClusterEngine, db.Status)
	}
}
