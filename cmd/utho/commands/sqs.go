package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sqsCmd = &cobra.Command{
	Use:   "sqs",
	Short: "Manage SQS instances",
}

var listSqsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all SQS instances",
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

		instances, err := client.SQS.List()
		if err != nil {
			fmt.Printf("Error listing SQS instances: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(instances, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(sqsCmd)
	sqsCmd.AddCommand(listSqsCmd)
}
