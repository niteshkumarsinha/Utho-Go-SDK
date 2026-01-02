package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/storage"
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
			fmt.Println("Error: API Key not configured.")
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

var createStorageCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new storage volume",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		size, _ := cmd.Flags().GetString("size")
		zone, _ := cmd.Flags().GetString("zone")

		params := storage.CreateParams{
			Name:   name,
			Size:   size,
			DCSlug: zone,
		}

		client, _ := utho.NewClient(apiKey)
		vol, err := client.Storage.Create(params)
		if err != nil {
			fmt.Printf("Error creating volume: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(vol, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteStorageCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a storage volume",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Storage.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting volume: %v\n", err)
			return
		}
		fmt.Println("Volume deleted successfully.")
	},
}

var attachStorageCmd = &cobra.Command{
	Use:   "attach [id]",
	Short: "Attach storage volume to a server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		serverID, _ := cmd.Flags().GetString("server")
		if serverID == "" {
			fmt.Println("Error: --server flag is required")
			return
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Storage.Attach(args[0], storage.AttachParams{ServerID: serverID})
		if err != nil {
			fmt.Printf("Error attaching volume: %v\n", err)
			return
		}
		fmt.Println("Volume attached successfully.")
	},
}

var detachStorageCmd = &cobra.Command{
	Use:   "detach [id]",
	Short: "Detach storage volume from a server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Storage.Detach(args[0])
		if err != nil {
			fmt.Printf("Error detaching volume: %v\n", err)
			return
		}
		fmt.Println("Volume detached successfully.")
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.AddCommand(listStorageCmd)
	storageCmd.AddCommand(createStorageCmd)
	storageCmd.AddCommand(deleteStorageCmd)
	storageCmd.AddCommand(attachStorageCmd)
	storageCmd.AddCommand(detachStorageCmd)

	createStorageCmd.Flags().String("name", "", "Name of the volume (required)")
	createStorageCmd.Flags().String("size", "", "Size in GB (required)")
	createStorageCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createStorageCmd.MarkFlagRequired("name")
	createStorageCmd.MarkFlagRequired("size")
	createStorageCmd.MarkFlagRequired("zone")

	attachStorageCmd.Flags().String("server", "", "Cloud Server ID (required)")
	attachStorageCmd.MarkFlagRequired("server")
}
