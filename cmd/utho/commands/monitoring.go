package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var monitoringCmd = &cobra.Command{
	Use:   "monitoring",
	Short: "Manage monitoring alert policies",
}

var listAlertPoliciesCmd = &cobra.Command{
	Use:   "alerts",
	Short: "List all alert policies",
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

		policies, err := client.Monitoring.ListAlertPolicies()
		if err != nil {
			fmt.Printf("Error listing alert policies: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(policies, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(monitoringCmd)
	monitoringCmd.AddCommand(listAlertPoliciesCmd)
}
