package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
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

func init() {
	rootCmd.AddCommand(kubernetesCmd)
	kubernetesCmd.AddCommand(listClustersCmd)
}
