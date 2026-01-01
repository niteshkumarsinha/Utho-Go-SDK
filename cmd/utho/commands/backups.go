package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
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

func init() {
	rootCmd.AddCommand(backupsCmd)
	backupsCmd.AddCommand(listBackupsCmd)
}
