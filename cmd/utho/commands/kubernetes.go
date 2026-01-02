package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/kubernetes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Manage Kubernetes clusters",
}

var listClustersCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Kubernetes clusters",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured. Run 'utho configure' or set UTHO_APIKEY environment variable.")
			return
		}

		client, err := utho.NewClient(apiKey)
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
			return
		}

		clusters, err := client.Kubernetes.List()
		if err != nil {
			fmt.Printf("Error listing clusters: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(clusters, "", "  ")
		fmt.Println(string(output))
	},
}

var getClusterCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get details of a Kubernetes cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		cluster, err := client.Kubernetes.Get(args[0])
		if err != nil {
			fmt.Printf("Error getting cluster: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(cluster, "", "  ")
		fmt.Println(string(output))
	},
}

var createClusterCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		zone, _ := cmd.Flags().GetString("zone")
		nodes, _ := cmd.Flags().GetInt("nodes")
		plan, _ := cmd.Flags().GetString("plan")

		params := kubernetes.CreateParams{
			ClusterLabel: name,
			DCSlug:       zone,
			NodePools: []kubernetes.NodePoolConfig{
				{
					Label: "default-pool",
					Count: nodes,
					Size:  plan,
				},
			},
		}

		client, _ := utho.NewClient(apiKey)
		cluster, err := client.Kubernetes.Create(params)
		if err != nil {
			fmt.Printf("Error creating cluster: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(cluster, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteClusterCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a Kubernetes cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Kubernetes.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting cluster: %v\n", err)
			return
		}
		fmt.Println("Cluster deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(kubernetesCmd)
	kubernetesCmd.AddCommand(listClustersCmd)
	kubernetesCmd.AddCommand(getClusterCmd)
	kubernetesCmd.AddCommand(createClusterCmd)
	kubernetesCmd.AddCommand(deleteClusterCmd)

	createClusterCmd.Flags().String("name", "", "Cluster Name (required)")
	createClusterCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createClusterCmd.Flags().Int("nodes", 1, "Number of worker nodes (required)")
	createClusterCmd.Flags().String("plan", "", "Worker Plan ID (required)")
	createClusterCmd.MarkFlagRequired("name")
	createClusterCmd.MarkFlagRequired("zone")
	createClusterCmd.MarkFlagRequired("nodes")
	createClusterCmd.MarkFlagRequired("plan")
}
