package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/networking"
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

var createDomainCmd = &cobra.Command{
	Use:   "create-domain [domain_name]",
	Short: "Add a new DNS domain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		params := networking.CreateDomainParams{
			Domain: args[0],
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Networking.CreateDomain(params)
		if err != nil {
			fmt.Printf("Error creating domain: %v\n", err)
			return
		}
		fmt.Println("Domain created successfully.")
	},
}

var deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain [domain_name]",
	Short: "Delete a DNS domain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Networking.DeleteDomain(args[0])
		if err != nil {
			fmt.Printf("Error deleting domain: %v\n", err)
			return
		}
		fmt.Println("Domain deleted successfully.")
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

var createFirewallCmd = &cobra.Command{
	Use:   "create-firewall [name]",
	Short: "Create a new firewall",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		params := networking.CreateFirewallParams{
			Name: args[0],
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Networking.CreateFirewall(params)
		if err != nil {
			fmt.Printf("Error creating firewall: %v\n", err)
			return
		}
		fmt.Println("Firewall created successfully.")
	},
}

var deleteFirewallCmd = &cobra.Command{
	Use:   "delete-firewall [id]",
	Short: "Delete a firewall",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		client, _ := utho.NewClient(apiKey)
		err := client.Networking.DeleteFirewall(args[0])
		if err != nil {
			fmt.Printf("Error deleting firewall: %v\n", err)
			return
		}
		fmt.Println("Firewall deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(networkingCmd)
	networkingCmd.AddCommand(listDomainsCmd)
	networkingCmd.AddCommand(createDomainCmd)
	networkingCmd.AddCommand(deleteDomainCmd)
	networkingCmd.AddCommand(listFirewallsCmd)
	networkingCmd.AddCommand(createFirewallCmd)
	networkingCmd.AddCommand(deleteFirewallCmd)
}
