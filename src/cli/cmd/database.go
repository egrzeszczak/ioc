package cmd

import (
	"github.com/spf13/cobra"
)

// ioc database
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Commands for database operations",
}

// ioc database audit
var databaseAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Review the last changes made to the iocdb database",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// ioc database import
var databaseImportCmd = &cobra.Command{
	Use:   "import <filepath>",
	Short: "Import a database from a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

// io database export
var databaseExportCmd = &cobra.Command{
	Use:   "export <filename>",
	Short: "Export the database to a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

func init() {
	databaseCmd.AddCommand(databaseAuditCmd)  // ioc database audit
	databaseCmd.AddCommand(databaseImportCmd) // ioc database import
	databaseCmd.AddCommand(databaseExportCmd) // ioc database export
	rootCmd.AddCommand(databaseCmd)           // ioc database
}
