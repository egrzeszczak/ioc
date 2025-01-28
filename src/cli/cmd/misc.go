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

// ioc list
var listCmd = &cobra.Command{
	Use:   "list <collection_name>",
	Short: "List all entries in a collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

// ioc add
var addCmd = &cobra.Command{
	Use:   "add <collection_name> <indicator>",
	Short: "Add an indicator to a collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}
	},
}

// ioc bulk-add
var bulkAddCmd = &cobra.Command{
	Use:   "bulk-add <collection_name> <filepath>",
	Short: "Add lots of indicators from a file to a collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}
	},
}

// ioc search
var searchCmd = &cobra.Command{
	Use:   "search <indicator>",
	Short: "Search for an indicator in all collections",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)    // ioc audit
	rootCmd.AddCommand(versionCmd)  // ioc version
	rootCmd.AddCommand(rollbackCmd) // ioc rollback
	rootCmd.AddCommand(listCmd)     // ioc list
	rootCmd.AddCommand(addCmd)      // ioc add
	rootCmd.AddCommand(bulkAddCmd)  // ioc bulk-add
	rootCmd.AddCommand(searchCmd)   // ioc search
}
