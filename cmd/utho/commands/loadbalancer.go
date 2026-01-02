package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/loadbalancer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "Manage load balancers",
}

var listLBsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all load balancers",
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

		lbs, err := client.LoadBalancer.List()
		if err != nil {
			fmt.Printf("Error listing load balancers: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(lbs, "", "  ")
		fmt.Println(string(output))
	},
}

var createLBCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new load balancer",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		zone, _ := cmd.Flags().GetString("zone")
		typ, _ := cmd.Flags().GetString("type")

		params := loadbalancer.CreateParams{
			Name:   name,
			DCSlug: zone,
			Type:   typ,
		}

		client, _ := utho.NewClient(apiKey)
		resp, err := client.LoadBalancer.Create(params)
		if err != nil {
			fmt.Printf("Error creating load balancer: %v\n", err)
			return
		}
		output, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(output))
	},
}

var deleteLBCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a load balancer",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.LoadBalancer.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting load balancer: %v\n", err)
			return
		}
		fmt.Println("Load balancer deleted successfully.")
	},
}

var updateLBCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a load balancer",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Error: One or more flags required for update")
			return
		}

		params := loadbalancer.UpdateParams{
			Name: name,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.LoadBalancer.Update(args[0], params)
		if err != nil {
			fmt.Printf("Error updating load balancer: %v\n", err)
			return
		}
		fmt.Println("Load balancer updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(loadbalancerCmd)
	loadbalancerCmd.AddCommand(listLBsCmd)
	loadbalancerCmd.AddCommand(createLBCmd)
	loadbalancerCmd.AddCommand(deleteLBCmd)
	loadbalancerCmd.AddCommand(updateLBCmd)

	createLBCmd.Flags().String("name", "", "Load Balancer Name (required)")
	createLBCmd.Flags().String("zone", "", "Zone/DC Slug (required)")
	createLBCmd.Flags().String("type", "public", "Type (public/private)")
	createLBCmd.MarkFlagRequired("name")
	createLBCmd.MarkFlagRequired("zone")

	updateLBCmd.Flags().String("name", "", "New Name")
}
