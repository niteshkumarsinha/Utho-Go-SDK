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

	fmt.Println("=== Kubernetes Example ===\n")

	// 1. List all Kubernetes clusters
	fmt.Println("1. Listing all Kubernetes clusters...")
	clusters, err := client.Kubernetes.List()
	if err != nil {
		log.Fatalf("Error listing clusters: %v", err)
	}
	fmt.Printf("Found %d clusters\n", len(clusters))
	for _, cluster := range clusters {
		fmt.Printf("  - ID: %s, Name: %s, Status: %s, Version: %s\n",
			cluster.ID, cluster.Name, cluster.Status, cluster.Version)
	}

	// 2. Get details of a specific cluster
	if len(clusters) > 0 {
		fmt.Println("\n2. Getting details of first cluster...")
		clusterID := clusters[0].ID
		cluster, err := client.Kubernetes.Get(clusterID)
		if err != nil {
			log.Printf("Error getting cluster: %v", err)
		} else {
			fmt.Printf("Cluster Details:\n")
			fmt.Printf("  ID: %s\n", cluster.ID)
			fmt.Printf("  Name: %s\n", cluster.Name)
			fmt.Printf("  Status: %s\n", cluster.Status)
			fmt.Printf("  Version: %s\n", cluster.Version)
		}
	}

	// 3. Create a new Kubernetes cluster (COMMENTED OUT)
	/*
		fmt.Println("\n3. Creating a new Kubernetes cluster...")
		createParams := kubernetes.CreateParams{
			ClusterLabel: "my-k8s-cluster",
			DCSlug:       "inmumbaizone2",
			NodePools: []kubernetes.NodePool{
				{
					Label: "worker-pool-1",
					Count: 2,
					Size:  "10045", // Plan ID
				},
			},
		}
		err = client.Kubernetes.Create(createParams)
		if err != nil {
			log.Printf("Error creating cluster: %v", err)
		} else {
			fmt.Println("Cluster creation initiated")
		}
	*/

	// 4. Delete a cluster (COMMENTED OUT)
	/*
		if len(clusters) > 0 {
			clusterID := clusters[0].ID
			fmt.Printf("\n4. Deleting cluster %s...\n", clusterID)
			err = client.Kubernetes.Delete(clusterID)
			if err != nil {
				log.Printf("Error deleting cluster: %v", err)
			} else {
				fmt.Println("Cluster deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
