package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
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

func init() {
	rootCmd.AddCommand(snapshotsCmd)
	snapshotsCmd.AddCommand(listSnapshotsCmd)
}
