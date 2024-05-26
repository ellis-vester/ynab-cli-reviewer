package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/brunomvsouza/ynab.go"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ellis-vester/ynab-cli-reviewer/tui"
	yn "github.com/ellis-vester/ynab-cli-reviewer/ynab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Categorize and approve your unapproved transactions",
	Run: func(cmd *cobra.Command, args []string) {

		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		configFileFolder := filepath.Join(homeDir, ".ynabr")

		viper.AddConfigPath(configFileFolder)
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		err = viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		apiKey := viper.GetString("apikey")

		client := yn.Ynab{
			Client: ynab.NewClient(apiKey),
		}

		model := tui.NewReviewModel(client)

		if _, err := tea.NewProgram(model).Run(); err != nil {
			fmt.Println("Oh no!!!!!", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}
