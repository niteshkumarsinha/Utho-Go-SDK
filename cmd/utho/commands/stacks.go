package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/stacks"
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

var createStackCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new stack",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		images, _ := cmd.Flags().GetString("images")
		script, _ := cmd.Flags().GetString("script")
		status, _ := cmd.Flags().GetString("status")

		params := stacks.CreateParams{
			Title:       title,
			Description: desc,
			Images:      images,
			Script:      script,
			Status:      status,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Stacks.Create(params)
		if err != nil {
			fmt.Printf("Error creating stack: %v\n", err)
			return
		}
		fmt.Println("Stack created successfully.")
	},
}

var deleteStackCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a stack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Stacks.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting stack: %v\n", err)
			return
		}
		fmt.Println("Stack deleted successfully.")
	},
}

var updateStackCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a stack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		images, _ := cmd.Flags().GetString("images")
		script, _ := cmd.Flags().GetString("script")
		status, _ := cmd.Flags().GetString("status")

		params := stacks.UpdateParams{
			Title:       title,
			Description: desc,
			Images:      images,
			Script:      script,
			Status:      status,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Stacks.Update(args[0], params)
		if err != nil {
			fmt.Printf("Error updating stack: %v\n", err)
			return
		}
		fmt.Println("Stack updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(stacksCmd)
	stacksCmd.AddCommand(listStacksCmd)
	stacksCmd.AddCommand(createStackCmd)
	stacksCmd.AddCommand(deleteStackCmd)
	stacksCmd.AddCommand(updateStackCmd)

	createStackCmd.Flags().String("title", "", "Stack Title (required)")
	createStackCmd.Flags().String("desc", "", "Description")
	createStackCmd.Flags().String("images", "", "Images (comma separated)")
	createStackCmd.Flags().String("script", "", "Script content")
	createStackCmd.Flags().String("status", "1", "Status (1=active)")
	createStackCmd.MarkFlagRequired("title")

	updateStackCmd.Flags().String("title", "", "Stack Title")
	updateStackCmd.Flags().String("desc", "", "Description")
	updateStackCmd.Flags().String("images", "", "Images")
	updateStackCmd.Flags().String("script", "", "Script content")
	updateStackCmd.Flags().String("status", "", "Status")
}
