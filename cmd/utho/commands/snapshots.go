package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/snapshots"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Manage cloud server snapshots",
}

var listSnapshotsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots",
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

		snapshots, err := client.Snapshots.List()
		if err != nil {
			fmt.Printf("Error listing snapshots: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(snapshots, "", "  ")
		fmt.Println(string(output))
	},
}

var createSnapshotCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		cloudID, _ := cmd.Flags().GetString("cloudid")
		name, _ := cmd.Flags().GetString("name")
		// Snapshot CreateParams does not have Description

		params := snapshots.CreateParams{
			CloudID: cloudID,
			Name:    name,
		}

		client, _ := utho.NewClient(apiKey)
		resp, err := client.Snapshots.Create(params)
		if err != nil {
			fmt.Printf("Error creating snapshot: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteSnapshotCmd = &cobra.Command{
	Use:   "delete [snapshot_id]",
	Short: "Delete a snapshot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		cloudID, _ := cmd.Flags().GetString("cloudid")
		if cloudID == "" {
			fmt.Println("Error: --cloudid flag is required for deletion.")
			return
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Snapshots.Delete(cloudID, args[0])
		if err != nil {
			fmt.Printf("Error deleting snapshot: %v\n", err)
			return
		}
		fmt.Println("Snapshot deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(snapshotsCmd)
	snapshotsCmd.AddCommand(listSnapshotsCmd)
	snapshotsCmd.AddCommand(createSnapshotCmd)
	snapshotsCmd.AddCommand(deleteSnapshotCmd)

	createSnapshotCmd.Flags().String("cloudid", "", "Cloud Server ID (required)")
	createSnapshotCmd.Flags().String("name", "", "Snapshot Name (required)")
	createSnapshotCmd.MarkFlagRequired("cloudid")
	createSnapshotCmd.MarkFlagRequired("name")

	deleteSnapshotCmd.Flags().String("cloudid", "", "Cloud Server ID (required)")
	deleteSnapshotCmd.MarkFlagRequired("cloudid")
}
