package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage container registries",
}

var listRegistriesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all container registries",
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

		registries, err := client.Registry.List()
		if err != nil {
			fmt.Printf("Error listing registries: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(registries, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
	registryCmd.AddCommand(listRegistriesCmd)
}
