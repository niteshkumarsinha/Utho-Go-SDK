package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/objectstorage"
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

var createBucketCmd = &cobra.Command{
	Use:   "create-bucket [dcslug]",
	Short: "Create a new object storage bucket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		dcslug := args[0]
		name, _ := cmd.Flags().GetString("name")
		size, _ := cmd.Flags().GetString("size")
		billing, _ := cmd.Flags().GetString("billing")

		params := objectstorage.CreateBucketParams{
			Name:    name,
			DCSlug:  dcslug,
			Size:    size,
			Billing: billing,
		}

		client, _ := utho.NewClient(apiKey)
		resp, err := client.ObjectStorage.CreateBucket(params)
		if err != nil {
			fmt.Printf("Error creating bucket: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteBucketCmd = &cobra.Command{
	Use:   "delete-bucket [dcslug] [name]",
	Short: "Delete an object storage bucket",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.ObjectStorage.DeleteBucket(args[0], args[1])
		if err != nil {
			fmt.Printf("Error deleting bucket: %v\n", err)
			return
		}
		fmt.Println("Bucket deleted successfully.")
	},
}

var listAccessKeysCmd = &cobra.Command{
	Use:   "list-keys [dcslug]",
	Short: "List access keys for a datacenter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		keys, err := client.ObjectStorage.ListAccessKeys(args[0])
		if err != nil {
			fmt.Printf("Error listing access keys: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(keys, "", "  ")
		fmt.Println(string(output))
	},
}

var createAccessKeyCmd = &cobra.Command{
	Use:   "create-key [dcslug]",
	Short: "Create a new access key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		key, err := client.ObjectStorage.CreateAccessKey(args[0])
		if err != nil {
			fmt.Printf("Error creating access key: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(key, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(objectstorageCmd)
	objectstorageCmd.AddCommand(listBucketsCmd)
	objectstorageCmd.AddCommand(createBucketCmd)
	objectstorageCmd.AddCommand(deleteBucketCmd)
	objectstorageCmd.AddCommand(listAccessKeysCmd)
	objectstorageCmd.AddCommand(createAccessKeyCmd)

	createBucketCmd.Flags().String("name", "", "Bucket Name (required)")
	createBucketCmd.Flags().String("size", "250GB", "Size")
	createBucketCmd.Flags().String("billing", "monthly", "Billing cycle")
	createBucketCmd.MarkFlagRequired("name")
}
