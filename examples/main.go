package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	// Initialize the client with your API key
	apiKey := os.Getenv("UTHO_API_KEY")
	if apiKey == "" {
		apiKey = "your-api-key"
	}

	client, err := utho.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	fmt.Println("=== Utho Go SDK - Full Service Demo ===")

	// 1. Compute & Management
	servers, _ := client.CloudServer.List()
	fmt.Printf("[CloudServer]  Found %d servers\n", len(servers))

	backups, _ := client.Backups.List()
	fmt.Printf("[Backups]      Found %d backups\n", len(backups))

	snaps, _ := client.Snapshots.List()
	fmt.Printf("[Snapshots]    Found %d snapshots\n", len(snaps))

	asGroups, _ := client.Autoscaling.List()
	fmt.Printf("[Autoscaling]  Found %d groups\n", len(asGroups))

	isos, _ := client.ISO.List()
	fmt.Printf("[ISO]          Found %d custom ISOs\n", len(isos))

	// 2. Managed Services
	clusters, _ := client.Kubernetes.List()
	fmt.Printf("[Kubernetes]   Found %d clusters\n", len(clusters))

	dbClusters, _ := client.Database.List()
	fmt.Printf("[Database]     Found %d clusters\n", len(dbClusters))

	buckets, _ := client.ObjectStorage.ListBuckets("inmumbaizone2")
	fmt.Printf("[ObjectStorage] Found %d buckets (Mumbai)\n", len(buckets))

	registries, _ := client.Registry.List()
	fmt.Printf("[Registry]     Found %d registries\n", len(registries))

	// 3. Networking
	vpcs, _ := client.VPC.List()
	fmt.Printf("[VPC]          Found %d VPCs\n", len(vpcs))

	lbs, _ := client.LoadBalancer.List()
	fmt.Printf("[LoadBalancer] Found %d load balancers\n", len(lbs))

	domains, _ := client.Networking.ListDomains()
	fmt.Printf("[DNS]          Found %d domains\n", len(domains))

	firewalls, _ := client.Networking.ListFirewalls()
	fmt.Printf("[Firewalls]   Found %d firewalls\n", len(firewalls))

	vpns, _ := client.VPN.List()
	fmt.Printf("[VPN]          Found %d VPNs\n", len(vpns))

	// 4. Storage
	ebs, _ := client.Storage.List()
	fmt.Printf("[EBS]          Found %d volumes\n", len(ebs))

	// 5. Security & Other
	keys, _ := client.Security.ListSSHKeys()
	fmt.Printf("[Security]     Found %d SSH keys\n", len(keys))

	wafs, _ := client.WAF.List()
	fmt.Printf("[WAF]          Found %d instances\n", len(wafs))

	certs, _ := client.SSL.List()
	fmt.Printf("[SSL]          Found %d certificates\n", len(certs))

	queues, _ := client.SQS.List()
	fmt.Printf("[SQS]          Found %d instances\n", len(queues))

	stacks, _ := client.Stacks.List()
	fmt.Printf("[Stacks]       Found %d templates\n", len(stacks))

	policies, _ := client.Monitoring.ListAlertPolicies()
	fmt.Printf("[Monitoring]   Found %d alert policies\n", len(policies))

	acc, _ := client.Account.GetInfo()
	if acc != nil {
		fmt.Printf("\n[Account] Balance: %s %s\n", acc.Balance, acc.Currency)
	}
}
