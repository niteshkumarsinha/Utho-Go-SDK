package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var networkingCmd = &cobra.Command{
	Use:   "networking",
	Short: "Manage DNS domains and firewalls",
}

var listDomainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "List all DNS domains",
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

		domains, err := client.Networking.ListDomains()
		if err != nil {
			fmt.Printf("Error listing domains: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(domains, "", "  ")
		fmt.Println(string(output))
	},
}

var listFirewallsCmd = &cobra.Command{
	Use:   "firewalls",
	Short: "List all firewalls",
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

		firewalls, err := client.Networking.ListFirewalls()
		if err != nil {
			fmt.Printf("Error listing firewalls: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(firewalls, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(networkingCmd)
	networkingCmd.AddCommand(listDomainsCmd)
	networkingCmd.AddCommand(listFirewallsCmd)
}
