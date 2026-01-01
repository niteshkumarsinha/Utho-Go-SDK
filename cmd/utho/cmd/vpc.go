package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
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

func init() {
	rootCmd.AddCommand(vpcCmd)
	vpcCmd.AddCommand(listVpcsCmd)
}
