package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "Manage WAF instances",
}

var listWafCmd = &cobra.Command{
	Use:   "list",
	Short: "List all WAF instances",
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

		wafs, err := client.WAF.List()
		if err != nil {
			fmt.Printf("Error listing WAFs: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(wafs, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(wafCmd)
	wafCmd.AddCommand(listWafCmd)
}
