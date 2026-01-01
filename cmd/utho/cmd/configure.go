package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure Utho CLI credentials",
	Run: func(cmd *cobra.Command, args []string) {
		var apiKey string
		fmt.Print("Utho API Key: ")
		fmt.Scanln(&apiKey)

		if apiKey == "" {
			fmt.Println("Error: API Key is required")
			return
		}

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		configDir := filepath.Join(home, ".utho")
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			os.MkdirAll(configDir, 0755)
		}

		config := map[string]string{
			"apikey": apiKey,
		}

		configPath := filepath.Join(configDir, "config.json")
		file, err := os.Create(configPath)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(config); err != nil {
			fmt.Println("Error writing config file:", err)
			return
		}

		fmt.Printf("Configuration saved to %s\n", configPath)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
