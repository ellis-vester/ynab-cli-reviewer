package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ellis-vester/ynab-cli-reviewer/config"
	"github.com/ellis-vester/ynab-cli-reviewer/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your API key for other commands to use.",
	Run: func(cmd *cobra.Command, args []string) {
		model := tui.NewInitModel()

		if _, err := tea.NewProgram(model).Run(); err != nil {
			fmt.Println("Oh no!!!", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configFileFolder := filepath.Join(homeDir, ".ynabr")
	configFilePath := filepath.Join(homeDir, ".ynabr", "config.json")

	_, err = os.Stat(configFilePath)
	if errors.Is(err, os.ErrNotExist) {

		fmt.Errorf("Could not find config file path, creating...")

		err := os.MkdirAll(configFileFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configFileFolder)

	// Set default config structure if not present
	var conf = config.Config{}

	viper.Set("config", conf)

	err = viper.SafeWriteConfig()
	if !strings.Contains(err.Error(), "Already Exists") {
		panic(err)
	}

	// Read the existing or just added default config into viper
	// for use elsewhere
	viper.ReadInConfig()
}
