package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/cloudserver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cloudserverCmd = &cobra.Command{
	Use:   "cloudserver",
	Short: "Manage cloud servers",
}

var listInstancesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all cloud server instances",
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

		instances, err := client.CloudServer.List()
		if err != nil {
			fmt.Printf("Error listing instances: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(instances, "", "  ")
		fmt.Println(string(output))
	},
}

var getInstanceCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get details of a specific cloud server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		instance, err := client.CloudServer.Get(args[0])
		if err != nil {
			fmt.Printf("Error getting instance: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(instance, "", "  ")
		fmt.Println(string(output))
	},
}

var createInstanceCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new cloud server",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		planID, _ := cmd.Flags().GetString("plan")
		zone, _ := cmd.Flags().GetString("zone")
		image, _ := cmd.Flags().GetString("image")
		// Add other flags as needed

		params := cloudserver.DeployParams{
			PlanID: planID,
			DCSlug: zone,
			Image:  image,
		}

		client, _ := utho.NewClient(apiKey)
		instance, err := client.CloudServer.Deploy(params)
		if err != nil {
			fmt.Printf("Error creating instance: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(instance, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteInstanceCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a cloud server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.CloudServer.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting instance: %v\n", err)
			return
		}
		fmt.Println("Instance deleted successfully.")
	},
}

var rebootInstanceCmd = &cobra.Command{
	Use:   "reboot [id]",
	Short: "Reboot a cloud server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.CloudServer.HardReboot(args[0])
		if err != nil {
			fmt.Printf("Error rebooting instance: %v\n", err)
			return
		}
		fmt.Println("Instance rebooted successfully.")
	},
}

var powerOnInstanceCmd = &cobra.Command{
	Use:   "power-on [id]",
	Short: "Power on a cloud server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.CloudServer.PowerOn(args[0])
		if err != nil {
			fmt.Printf("Error powering on instance: %v\n", err)
			return
		}
		fmt.Println("Instance powered on successfully.")
	},
}

var powerOffInstanceCmd = &cobra.Command{
	Use:   "power-off [id]",
	Short: "Power off a cloud server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.CloudServer.PowerOff(args[0])
		if err != nil {
			fmt.Printf("Error powering off instance: %v\n", err)
			return
		}
		fmt.Println("Instance powered off successfully.")
	},
}

var resetPasswordCmd = &cobra.Command{
	Use:   "reset-password [id]",
	Short: "Reset cloud server root password",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.CloudServer.ResetPassword(args[0])
		if err != nil {
			fmt.Printf("Error resetting password: %v\n", err)
			return
		}
		fmt.Println("Password reset successfully. Check your email for new credentials.")
	},
}

func init() {
	rootCmd.AddCommand(cloudserverCmd)
	cloudserverCmd.AddCommand(listInstancesCmd)
	cloudserverCmd.AddCommand(getInstanceCmd)
	cloudserverCmd.AddCommand(createInstanceCmd)
	cloudserverCmd.AddCommand(deleteInstanceCmd)
	cloudserverCmd.AddCommand(rebootInstanceCmd)
	cloudserverCmd.AddCommand(powerOnInstanceCmd)
	cloudserverCmd.AddCommand(powerOffInstanceCmd)
	cloudserverCmd.AddCommand(resetPasswordCmd)

	createInstanceCmd.Flags().String("plan", "", "Plan ID (required)")
	createInstanceCmd.Flags().String("zone", "", "Zone (required)")
	createInstanceCmd.Flags().String("image", "", "Image (required)")
	createInstanceCmd.MarkFlagRequired("plan")
	createInstanceCmd.MarkFlagRequired("zone")
	createInstanceCmd.MarkFlagRequired("image")
}
