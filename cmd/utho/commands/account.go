package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage account information",
}

var getAccountInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get account information",
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

		info, err := client.Account.GetInfo()
		if err != nil {
			fmt.Printf("Error getting account info: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(getAccountInfoCmd)
}
