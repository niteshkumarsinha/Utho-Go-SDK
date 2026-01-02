package commands

import (
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/transfer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Initiate and manage resource transfers",
}

var initiateTransferCmd = &cobra.Command{
	Use:   "initiate [type] [id]",
	Short: "Initiate a resource transfer",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured. Run 'utho configure' or set UTHO_APIKEY environment variable.")
			return
		}

		resourceType := args[0]
		resourceID := args[1]

		client, err := utho.NewClient(apiKey)
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
			return
		}

		err = client.Transfer.Initiate(resourceType, resourceID)
		if err != nil {
			fmt.Printf("Error initiating transfer: %v\n", err)
			return
		}

		fmt.Println("Transfer initiated successfully.")
	},
}

var receiveTransferCmd = &cobra.Command{
	Use:   "receive [type] [id] [token]",
	Short: "Receive a transferred resource",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		resourceType := args[0]
		resourceID := args[1]
		token := args[2]

		params := transfer.ReceiveParams{
			ID:    resourceID,
			Token: token,
			Type:  resourceType,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Transfer.Receive(params)
		if err != nil {
			fmt.Printf("Error receiving transfer: %v\n", err)
			return
		}
		fmt.Println("Resource transfer received successfully.")
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	transferCmd.AddCommand(initiateTransferCmd)
	transferCmd.AddCommand(receiveTransferCmd)
}
