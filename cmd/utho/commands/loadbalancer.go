package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "Manage load balancers",
}

var listLBsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all load balancers",
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

		lbs, err := client.LoadBalancer.List()
		if err != nil {
			fmt.Printf("Error listing load balancers: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(lbs, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(loadbalancerCmd)
	loadbalancerCmd.AddCommand(listLBsCmd)
}
