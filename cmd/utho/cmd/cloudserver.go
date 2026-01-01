package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cloudserverCmd = &cobra.Command{
	Use:   "cloudserver",
	Short: "Manage cloud servers",
}

var listInstancesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all cloud server instances",
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

		instances, err := client.CloudServer.List()
		if err != nil {
			fmt.Printf("Error listing instances: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(instances, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(cloudserverCmd)
	cloudserverCmd.AddCommand(listInstancesCmd)
}
