package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stacksCmd = &cobra.Command{
	Use:   "stacks",
	Short: "Manage automation stacks",
}

var listStacksCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stacks",
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

		stacks, err := client.Stacks.List()
		if err != nil {
			fmt.Printf("Error listing stacks: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(stacks, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(stacksCmd)
	stacksCmd.AddCommand(listStacksCmd)
}
