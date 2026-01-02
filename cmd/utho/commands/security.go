package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/security"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var securityCmd = &cobra.Command{
	Use:   "security",
	Short: "Manage SSH and API keys",
}

var listSSHKeysCmd = &cobra.Command{
	Use:   "list-ssh",
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

var importSSHKeyCmd = &cobra.Command{
	Use:   "import-ssh",
	Short: "Import a new SSH key",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		key, _ := cmd.Flags().GetString("key")

		params := security.ImportSSHKeyParams{
			Name:      name,
			PublicKey: key,
		}

		client, _ := utho.NewClient(apiKey)
		resp, err := client.Security.ImportSSHKey(params)
		if err != nil {
			fmt.Printf("Error importing SSH key: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteSSHKeyCmd = &cobra.Command{
	Use:   "delete-ssh [id]",
	Short: "Delete an SSH key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Security.DeleteSSHKey(args[0])
		if err != nil {
			fmt.Printf("Error deleting SSH key: %v\n", err)
			return
		}
		fmt.Println("SSH key deleted successfully.")
	},
}

var listAPIKeysCmd = &cobra.Command{
	Use:   "list-api",
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

var generateAPIKeyCmd = &cobra.Command{
	Use:   "generate-api",
	Short: "Generate a new API key",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		label, _ := cmd.Flags().GetString("label")

		params := security.GenerateAPIKeyParams{
			Label: label,
		}

		client, _ := utho.NewClient(apiKey)
		key, err := client.Security.GenerateAPIKey(params)
		if err != nil {
			fmt.Printf("Error generating API key: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(key, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteAPIKeyCmd = &cobra.Command{
	Use:   "delete-api [id]",
	Short: "Delete an API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Security.DeleteAPIKey(args[0])
		if err != nil {
			fmt.Printf("Error deleting API key: %v\n", err)
			return
		}
		fmt.Println("API key deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(securityCmd)
	securityCmd.AddCommand(listSSHKeysCmd)
	securityCmd.AddCommand(importSSHKeyCmd)
	securityCmd.AddCommand(deleteSSHKeyCmd)
	securityCmd.AddCommand(listAPIKeysCmd)
	securityCmd.AddCommand(generateAPIKeyCmd)
	securityCmd.AddCommand(deleteAPIKeyCmd)

	importSSHKeyCmd.Flags().String("name", "", "Key Name (required)")
	importSSHKeyCmd.Flags().String("key", "", "Public Key string (required)")
	importSSHKeyCmd.MarkFlagRequired("name")
	importSSHKeyCmd.MarkFlagRequired("key")

	generateAPIKeyCmd.Flags().String("label", "", "API Key Label (required)")
	generateAPIKeyCmd.MarkFlagRequired("label")
}
