package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/backups"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var backupsCmd = &cobra.Command{
	Use:   "backups",
	Short: "Manage cloud server backups",
}

var listBackupsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all backups",
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

		backups, err := client.Backups.List()
		if err != nil {
			fmt.Printf("Error listing backups: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(backups, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteBackupCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a backup",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Backups.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting backup: %v\n", err)
			return
		}
		fmt.Println("Backup deleted successfully.")
	},
}

var restoreBackupCmd = &cobra.Command{
	Use:   "restore [id]",
	Short: "Restore a backup to a server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		cloudID, _ := cmd.Flags().GetString("cloudid")

		client, _ := utho.NewClient(apiKey)
		err := client.Backups.Restore(args[0], backups.RestoreParams{CloudID: cloudID})
		if err != nil {
			fmt.Printf("Error restoring backup: %v\n", err)
			return
		}
		fmt.Println("Backup restoration initiated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(backupsCmd)
	backupsCmd.AddCommand(listBackupsCmd)
	backupsCmd.AddCommand(deleteBackupCmd)
	backupsCmd.AddCommand(restoreBackupCmd)

	restoreBackupCmd.Flags().String("cloudid", "", "Target Cloud Server ID (required)")
	restoreBackupCmd.MarkFlagRequired("cloudid")
}
