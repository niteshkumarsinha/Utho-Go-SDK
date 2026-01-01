package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var isoCmd = &cobra.Command{
	Use:   "iso",
	Short: "Manage custom ISO images",
}

var listIsoCmd = &cobra.Command{
	Use:   "list",
	Short: "List all custom ISOs",
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

		isos, err := client.ISO.List()
		if err != nil {
			fmt.Printf("Error listing ISOs: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(isos, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(isoCmd)
	isoCmd.AddCommand(listIsoCmd)
}
