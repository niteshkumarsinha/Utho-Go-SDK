package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/monitoring"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var monitoringCmd = &cobra.Command{
	Use:   "monitoring",
	Short: "Manage monitoring alert policies",
}

var listAlertPoliciesCmd = &cobra.Command{
	Use:   "alerts",
	Short: "List all alert policies",
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

		policies, err := client.Monitoring.ListAlertPolicies()
		if err != nil {
			fmt.Printf("Error listing alert policies: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(policies, "", "  ")
		fmt.Println(string(output))
	},
}

var createAlertPolicyCmd = &cobra.Command{
	Use:   "create-alert",
	Short: "Create a new alert policy",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		label, _ := cmd.Flags().GetString("label")
		resType, _ := cmd.Flags().GetString("resource-type")
		contacts, _ := cmd.Flags().GetStringSlice("contacts")
		cpu, _ := cmd.Flags().GetInt("cpu")
		ram, _ := cmd.Flags().GetInt("ram")
		disk, _ := cmd.Flags().GetInt("disk")
		bw, _ := cmd.Flags().GetInt("bandwidth")

		params := monitoring.CreateAlertPolicyParams{
			Label:        label,
			ResourceType: resType,
			Contacts:     contacts,
		}
		params.Thresholds.CPU = cpu
		params.Thresholds.RAM = ram
		params.Thresholds.Disk = disk
		params.Thresholds.Bandwidth = bw

		client, _ := utho.NewClient(apiKey)
		err := client.Monitoring.CreateAlertPolicy(params)
		if err != nil {
			fmt.Printf("Error creating alert policy: %v\n", err)
			return
		}
		fmt.Println("Alert policy created successfully.")
	},
}

var deleteAlertPolicyCmd = &cobra.Command{
	Use:   "delete-alert [id]",
	Short: "Delete an alert policy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.Monitoring.DeleteAlertPolicy(args[0])
		if err != nil {
			fmt.Printf("Error deleting alert policy: %v\n", err)
			return
		}
		fmt.Println("Alert policy deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(monitoringCmd)
	monitoringCmd.AddCommand(listAlertPoliciesCmd)
	monitoringCmd.AddCommand(createAlertPolicyCmd)
	monitoringCmd.AddCommand(deleteAlertPolicyCmd)

	createAlertPolicyCmd.Flags().String("label", "", "Policy Label (required)")
	createAlertPolicyCmd.Flags().String("resource-type", "cloudserver", "Resource Type")
	createAlertPolicyCmd.Flags().StringSlice("contacts", []string{}, "Contact IDs")
	createAlertPolicyCmd.Flags().Int("cpu", 0, "CPU Threshold %")
	createAlertPolicyCmd.Flags().Int("ram", 0, "RAM Threshold %")
	createAlertPolicyCmd.Flags().Int("disk", 0, "Disk Threshold %")
	createAlertPolicyCmd.Flags().Int("bandwidth", 0, "Bandwidth Threshold %")
	createAlertPolicyCmd.MarkFlagRequired("label")
}
