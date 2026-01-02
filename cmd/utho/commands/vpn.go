package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/vpn"
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

var createVpnCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new VPN instance",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		zone, _ := cmd.Flags().GetString("zone")
		plan, _ := cmd.Flags().GetString("plan")

		params := vpn.CreateParams{
			Name:   name,
			DCSlug: zone,
			Plan:   plan,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.VPN.Create(params)
		if err != nil {
			fmt.Printf("Error creating VPN: %v\n", err)
			return
		}
		fmt.Println("VPN created successfully.")
	},
}

var deleteVpnCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a VPN instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.VPN.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting VPN: %v\n", err)
			return
		}
		fmt.Println("VPN deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(vpnCmd)
	vpnCmd.AddCommand(listVpnCmd)
	vpnCmd.AddCommand(createVpnCmd)
	vpnCmd.AddCommand(deleteVpnCmd)

	createVpnCmd.Flags().String("name", "", "VPN Name (required)")
	createVpnCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createVpnCmd.Flags().String("plan", "", "Plan ID (required)")
	createVpnCmd.MarkFlagRequired("name")
	createVpnCmd.MarkFlagRequired("zone")
	createVpnCmd.MarkFlagRequired("plan")
}
