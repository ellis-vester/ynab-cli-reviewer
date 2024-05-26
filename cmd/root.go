package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "ynabr",
	Short:   "A tool for reviewing your pending YNAB transactions.",
	Version: Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

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

	err = viper.SafeWriteConfig()
	if err != nil {
		if !strings.Contains(err.Error(), "Already Exists") {
			panic(err)
		}
	}
}
