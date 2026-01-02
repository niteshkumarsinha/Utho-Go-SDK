package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/autoscaling"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var autoscalingCmd = &cobra.Command{
	Use:   "autoscaling",
	Short: "Manage autoscaling groups",
}

var listASGroupsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all autoscaling groups",
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

		groups, err := client.Autoscaling.List()
		if err != nil {
			fmt.Printf("Error listing groups: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(groups, "", "  ")
		fmt.Println(string(output))
	},
}

var createASGroupCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new autoscaling group",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		min, _ := cmd.Flags().GetInt("min")
		max, _ := cmd.Flags().GetInt("max")
		image, _ := cmd.Flags().GetString("image")
		plan, _ := cmd.Flags().GetString("plan")
		script, _ := cmd.Flags().GetString("script")

		params := autoscaling.CreateParams{
			Name:    name,
			MinSize: min,
			MaxSize: max,
			Image:   image,
			Plan:    plan,
			Script:  script,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Autoscaling.Create(params)
		if err != nil {
			fmt.Printf("Error creating autoscaling group: %v\n", err)
			return
		}
		fmt.Println("Autoscaling group created successfully.")
	},
}

var deleteASGroupCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete an autoscaling group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Autoscaling.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting group: %v\n", err)
			return
		}
		fmt.Println("Autoscaling group deleted successfully.")
	},
}

var updateASGroupCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an autoscaling group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		min, _ := cmd.Flags().GetInt("min")
		max, _ := cmd.Flags().GetInt("max")

		params := autoscaling.UpdateParams{
			MinSize: min,
			MaxSize: max,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Autoscaling.Update(args[0], params)
		if err != nil {
			fmt.Printf("Error updating group: %v\n", err)
			return
		}
		fmt.Println("Autoscaling group updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(autoscalingCmd)
	autoscalingCmd.AddCommand(listASGroupsCmd)
	autoscalingCmd.AddCommand(createASGroupCmd)
	autoscalingCmd.AddCommand(deleteASGroupCmd)
	autoscalingCmd.AddCommand(updateASGroupCmd)

	createASGroupCmd.Flags().String("name", "", "Group Name (required)")
	createASGroupCmd.Flags().Int("min", 1, "Min Size (default: 1)")
	createASGroupCmd.Flags().Int("max", 2, "Max Size (default: 2)")
	createASGroupCmd.Flags().String("image", "", "Image slug (required)")
	createASGroupCmd.Flags().String("plan", "", "Plan ID (required)")
	createASGroupCmd.Flags().String("script", "", "Startup script")
	createASGroupCmd.MarkFlagRequired("name")
	createASGroupCmd.MarkFlagRequired("image")
	createASGroupCmd.MarkFlagRequired("plan")

	updateASGroupCmd.Flags().Int("min", 1, "Min Size")
	updateASGroupCmd.Flags().Int("max", 2, "Max Size")
}
