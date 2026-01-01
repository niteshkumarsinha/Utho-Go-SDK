package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var autoscalingCmd = &cobra.Command{
	Use:   "autoscaling",
	Short: "Manage autoscaling groups",
}

var listASGroupsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all autoscaling groups",
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

		groups, err := client.Autoscaling.List()
		if err != nil {
			fmt.Printf("Error listing groups: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(groups, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(autoscalingCmd)
	autoscalingCmd.AddCommand(listASGroupsCmd)
}
