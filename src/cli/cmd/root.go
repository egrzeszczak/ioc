package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// ioc
var rootCmd = &cobra.Command{
	Use:   "ioc",
	Short: "Indicator of Compromise CLI management tool for Linux",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true, // Disable the `completion` command
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
