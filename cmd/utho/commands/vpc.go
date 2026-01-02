package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/vpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Manage VPCs",
}

var listVpcsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all VPCs",
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

		vpcs, err := client.VPC.List()
		if err != nil {
			fmt.Printf("Error listing VPCs: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(vpcs, "", "  ")
		fmt.Println(string(output))
	},
}

var createVpcCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new VPC",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		region, _ := cmd.Flags().GetString("region")
		ipRange, _ := cmd.Flags().GetString("range")

		params := vpc.CreateParams{
			Name:      name,
			Region:    region,
			IPv4Range: ipRange,
		}

		client, _ := utho.NewClient(apiKey)
		resp, err := client.VPC.Create(params)
		if err != nil {
			fmt.Printf("Error creating VPC: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteVpcCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a VPC",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.VPC.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting VPC: %v\n", err)
			return
		}
		fmt.Println("VPC deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(vpcCmd)
	vpcCmd.AddCommand(listVpcsCmd)
	vpcCmd.AddCommand(createVpcCmd)
	vpcCmd.AddCommand(deleteVpcCmd)

	createVpcCmd.Flags().String("name", "", "VPC Name (required)")
	createVpcCmd.Flags().String("region", "", "Region (required)")
	createVpcCmd.Flags().String("range", "", "IPv4 Range (e.g., 10.0.0.0/16) (required)")
	createVpcCmd.MarkFlagRequired("name")
	createVpcCmd.MarkFlagRequired("region")
	createVpcCmd.MarkFlagRequired("range")
}
