package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "utho",
	Short: "Utho CLI - Manage your Utho cloud resources",
	Long:  `A command line interface for interacting with the Utho Cloud API.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.utho/config.json)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigType("json")
		viper.AddConfigPath(filepath.Join(home, ".utho"))
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("UTHO")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// Config file found and successfully read
	}
}
