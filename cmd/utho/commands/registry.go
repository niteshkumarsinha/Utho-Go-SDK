package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/registry"
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

var createRegistryCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new container registry",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		zone, _ := cmd.Flags().GetString("zone")

		params := registry.CreateParams{
			Name:   name,
			DCSlug: zone,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Registry.Create(params)
		if err != nil {
			fmt.Printf("Error creating registry: %v\n", err)
			return
		}
		fmt.Println("Registry created successfully.")
	},
}

var deleteRegistryCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a container registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Registry.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting registry: %v\n", err)
			return
		}
		fmt.Println("Registry deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
	registryCmd.AddCommand(listRegistriesCmd)
	registryCmd.AddCommand(createRegistryCmd)
	registryCmd.AddCommand(deleteRegistryCmd)

	createRegistryCmd.Flags().String("name", "", "Registry Name (required)")
	createRegistryCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createRegistryCmd.MarkFlagRequired("name")
	createRegistryCmd.MarkFlagRequired("zone")
}
