package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
