package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Manage database clusters",
}

var listDatabasesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all database clusters",
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

		databases, err := client.Database.List()
		if err != nil {
			fmt.Printf("Error listing databases: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(databases, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(listDatabasesCmd)
}
