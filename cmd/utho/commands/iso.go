package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/iso"
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

var createIsoCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new custom ISO from URL",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		url, _ := cmd.Flags().GetString("url")
		zone, _ := cmd.Flags().GetString("zone")

		params := iso.CreateParams{
			Name:   name,
			URL:    url,
			DCSlug: zone,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.ISO.Create(params)
		if err != nil {
			fmt.Printf("Error creating ISO: %v\n", err)
			return
		}
		fmt.Println("ISO added successfully.")
	},
}

var deleteIsoCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a custom ISO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.ISO.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting ISO: %v\n", err)
			return
		}
		fmt.Println("ISO deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(isoCmd)
	isoCmd.AddCommand(listIsoCmd)
	isoCmd.AddCommand(createIsoCmd)
	isoCmd.AddCommand(deleteIsoCmd)

	createIsoCmd.Flags().String("name", "", "ISO Name (required)")
	createIsoCmd.Flags().String("url", "", "ISO URL (required)")
	createIsoCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createIsoCmd.MarkFlagRequired("name")
	createIsoCmd.MarkFlagRequired("url")
	createIsoCmd.MarkFlagRequired("zone")
}
