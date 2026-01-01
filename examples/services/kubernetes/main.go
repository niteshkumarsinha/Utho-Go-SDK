package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/kubernetes"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	fmt.Println("Creating a Kubernetes cluster...")

	params := kubernetes.CreateParams{
		DCSlug:         "inmumbaizone2",
		ClusterLabel:   "prod-cluster",
		ClusterVersion: "v1.28",
		NodePools: []kubernetes.NodePoolConfig{
			{Label: "worker-pool", Size: "10045", Count: 3},
		},
	}

	resp, err := client.Kubernetes.Create(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Cluster creation started: %s\n", resp.ClusterID)

	clusters, _ := client.Kubernetes.List()
	for _, c := range clusters {
		fmt.Printf("- %s: %s [%s]\n", c.ClusterLabel, c.DC, c.Status)
	}
}
