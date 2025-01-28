package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ioc version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ioc@v1.0.0, github.com/egrzeszczak/ioc")
	},
}

// ioc audit
var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Review the last changes made to the iocdb sdatabase",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// ioc rollback
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback the last change made to the iocdb database",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)    // ioc audit
	rootCmd.AddCommand(versionCmd)  // ioc version
	rootCmd.AddCommand(rollbackCmd) // ioc rollback
}
