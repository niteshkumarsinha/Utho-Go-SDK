package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sslCmd = &cobra.Command{
	Use:   "ssl",
	Short: "Manage SSL certificates",
}

var listSslCmd = &cobra.Command{
	Use:   "list",
	Short: "List all SSL certificates",
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

		certs, err := client.SSL.List()
		if err != nil {
			fmt.Printf("Error listing certificates: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(certs, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(sslCmd)
	sslCmd.AddCommand(listSslCmd)
}
