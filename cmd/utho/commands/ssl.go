package commands

import (
	"encoding/json"
	"fmt"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/ssl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sslCmd = &cobra.Command{
	Use:   "ssl",
	Short: "Manage SSL certificates",
}

var listSslCmd = &cobra.Command{
	Use:   "list",
	Short: "List all SSL certificates",
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

		certs, err := client.SSL.List()
		if err != nil {
			fmt.Printf("Error listing certificates: %v\n", err)
			return
		}

		output, _ := json.MarshalIndent(certs, "", "  ")
		fmt.Println(string(output))
	},
}

var createSslCmd = &cobra.Command{
	Use:   "create",
	Short: "Upload a new SSL certificate",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		cert, _ := cmd.Flags().GetString("cert")
		key, _ := cmd.Flags().GetString("key")
		chain, _ := cmd.Flags().GetString("chain")

		params := ssl.CreateParams{
			Name:        name,
			Certificate: cert,
			PrivateKey:  key,
			Chain:       chain,
		}

		client, _ := utho.NewClient(apiKey)
		err := client.SSL.Create(params)
		if err != nil {
			fmt.Printf("Error uploading certificate: %v\n", err)
			return
		}
		fmt.Println("Certificate uploaded successfully.")
	},
}

var deleteSslCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete an SSL certificate",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("apikey")
		if apiKey == "" {
			fmt.Println("Error: API Key not configured.")
			return
		}
		client, _ := utho.NewClient(apiKey)
		err := client.SSL.Delete(args[0])
		if err != nil {
			fmt.Printf("Error deleting certificate: %v\n", err)
			return
		}
		fmt.Println("Certificate deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(sslCmd)
	sslCmd.AddCommand(listSslCmd)
	sslCmd.AddCommand(createSslCmd)
	sslCmd.AddCommand(deleteSslCmd)

	createSslCmd.Flags().String("name", "", "Certificate Name (required)")
	createSslCmd.Flags().String("cert", "", "Certificate body (required)")
	createSslCmd.Flags().String("key", "", "Private Key body (required)")
	createSslCmd.Flags().String("chain", "", "CA Chain body")
	createSslCmd.MarkFlagRequired("name")
	createSslCmd.MarkFlagRequired("cert")
	createSslCmd.MarkFlagRequired("key")
}
