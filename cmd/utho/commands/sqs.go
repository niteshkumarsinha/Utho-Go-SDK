package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/sqs"
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

var createSqsCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new SQS instance",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		zone, _ := cmd.Flags().GetString("zone")

		params := sqs.CreateParams{
			Name:   name,
			DCSlug: zone,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.SQS.Create(params)
		if err != nil {
			fmt.Printf("Error creating SQS instance: %v\n", err)
			return
		}
		fmt.Println("SQS instance created successfully.")
	},
}

var deleteSqsCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete an SQS instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.SQS.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting SQS instance: %v\n", err)
			return
		}
		fmt.Println("SQS instance deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(sqsCmd)
	sqsCmd.AddCommand(listSqsCmd)
	sqsCmd.AddCommand(createSqsCmd)
	sqsCmd.AddCommand(deleteSqsCmd)

	createSqsCmd.Flags().String("name", "", "SQS Name (required)")
	createSqsCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createSqsCmd.MarkFlagRequired("name")
	createSqsCmd.MarkFlagRequired("zone")
}
