package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage storage volumes",
}

var listStorageCmd = &cobra.Command{
	Use:   "list",
	Short: "List all storage volumes",
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

		volumes, err := client.Storage.List()
		if err != nil {
			fmt.Printf("Error listing storage volumes: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(volumes, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.AddCommand(listStorageCmd)
}
