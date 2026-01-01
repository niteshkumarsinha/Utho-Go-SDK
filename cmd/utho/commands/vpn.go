package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "Manage VPN instances",
}

var listVpnCmd = &cobra.Command{
	Use:   "list",
	Short: "List all VPN instances",
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

		vpns, err := client.VPN.List()
		if err != nil {
			fmt.Printf("Error listing VPNs: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(vpns, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(vpnCmd)
	vpnCmd.AddCommand(listVpnCmd)
}
