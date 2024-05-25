package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ellis-vester/ynab-cli-reviewer/tui"
	"github.com/spf13/cobra"
)

var Version string = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "ynabr",
	Short:   "A tool for reviewing your pending YNAB transactions.",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		model := tui.NewReviewModel()

		if _, err := tea.NewProgram(model).Run(); err != nil {
			fmt.Println("Oh no!!!", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
