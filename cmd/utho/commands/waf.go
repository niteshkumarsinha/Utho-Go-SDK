package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/waf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "Manage WAF instances",
}

var listWafCmd = &cobra.Command{
	Use:   "list",
	Short: "List all WAF instances",
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

		wafs, err := client.WAF.List()
		if err != nil {
			fmt.Printf("Error listing WAFs: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(wafs, "", "  ")
		fmt.Println(string(output))
	},
}

var createWafCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new WAF instance",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		zone, _ := cmd.Flags().GetString("zone")

		params := waf.CreateParams{
			Name:   name,
			DCSlug: zone,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.WAF.Create(params)
		if err != nil {
			fmt.Printf("Error creating WAF: %v\n", err)
			return
		}
		fmt.Println("WAF created successfully.")
	},
}

var deleteWafCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a WAF instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.WAF.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting WAF: %v\n", err)
			return
		}
		fmt.Println("WAF deleted successfully.")
	},
}

var attachWafCmd = &cobra.Command{
	Use:   "attach [id]",
	Short: "Attach WAF to a resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		resID, _ := cmd.Flags().GetString("resource-id")
		resType, _ := cmd.Flags().GetString("resource-type")

		params := waf.AttachParams{
			ResourceID:   resID,
			ResourceType: resType,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.WAF.Attach(args[0], params)
		if err != nil {
			fmt.Printf("Error attaching WAF: %v\n", err)
			return
		}
		fmt.Println("WAF attached successfully.")
	},
}

var detachWafCmd = &cobra.Command{
	Use:   "detach [id]",
	Short: "Detach WAF from a resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		client, _ := utho.NewClient(apiKey)
		err := client.WAF.Detach(args[0])
		if err != nil {
			fmt.Printf("Error detaching WAF: %v\n", err)
			return
		}
		fmt.Println("WAF detached successfully.")
	},
}

func init() {
	rootCmd.AddCommand(wafCmd)
	wafCmd.AddCommand(listWafCmd)
	wafCmd.AddCommand(createWafCmd)
	wafCmd.AddCommand(deleteWafCmd)
	wafCmd.AddCommand(attachWafCmd)
	wafCmd.AddCommand(detachWafCmd)

	createWafCmd.Flags().String("name", "", "WAF Name (required)")
	createWafCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createWafCmd.MarkFlagRequired("name")
	createWafCmd.MarkFlagRequired("zone")

	attachWafCmd.Flags().String("resource-id", "", "Resource ID to attach to (required)")
	attachWafCmd.Flags().String("resource-type", "loadbalancer", "Resource Type (default: loadbalancer)")
	attachWafCmd.MarkFlagRequired("resource-id")
}
