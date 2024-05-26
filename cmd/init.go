package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ellis-vester/ynab-cli-reviewer/tui"
	"github.com/spf13/cobra"
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
}
