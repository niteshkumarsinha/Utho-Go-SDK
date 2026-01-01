package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var objectstorageCmd = &cobra.Command{
	Use:   "objectstorage",
	Short: "Manage object storage buckets",
}

var listBucketsCmd = &cobra.Command{
	Use:   "list [dcslug]",
	Short: "List all buckets in a specific datacenter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured. Run 'utho configure' or set UTHO_APIKEY environment variable.")
			return
		}

		dcslug := args[0]
		client, err := utho.NewClient(apiKey)
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
			return
		}

		buckets, err := client.ObjectStorage.ListBuckets(dcslug)
		if err != nil {
			fmt.Printf("Error listing buckets: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(buckets, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(objectstorageCmd)
	objectstorageCmd.AddCommand(listBucketsCmd)
}
