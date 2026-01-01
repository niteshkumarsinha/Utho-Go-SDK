package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var securityCmd = &cobra.Command{
	Use:   "security",
	Short: "Manage SSH and API keys",
}

var listSSHKeysCmd = &cobra.Command{
	Use:   "sshkeys",
	Short: "List all SSH keys",
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

		keys, err := client.Security.ListSSHKeys()
		if err != nil {
			fmt.Printf("Error listing SSH keys: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(keys, "", "  ")
		fmt.Println(string(output))
	},
}

var listAPIKeysCmd = &cobra.Command{
	Use:   "apikeys",
	Short: "List all API keys",
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

		keys, err := client.Security.ListAPIKeys()
		if err != nil {
			fmt.Printf("Error listing API keys: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(keys, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(securityCmd)
	securityCmd.AddCommand(listSSHKeysCmd)
	securityCmd.AddCommand(listAPIKeysCmd)
}
