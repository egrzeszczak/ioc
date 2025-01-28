package cmd

import (
	"fmt"

	"github.com/egrzeszczak/ioc/src/cli/functions"
	"github.com/spf13/cobra"
)

// ioc collection
var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Create, modify and manage collections",
	Long:  "Create, modify and manage collections. Collections are containers for indicators.",
}

// ioc collection list
var collectionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all collections",
	Run: func(cmd *cobra.Command, args []string) {
		collections, err := functions.GetCollections()
		if err != nil {
			fmt.Printf("Error getting collections: %v\n", err)
			return
		}

		// Print collections to the console: id\tname\twhitelist
		fmt.Printf("ID\t\tName\n")
		for _, c := range collections {
			fmt.Printf("%d\t\t%s\n", c.ID, c.Name)
		}
	},
}

// ioc collection create
var collectionCreateCmd = &cobra.Command{
	Use:   "create <collection_name>",
	Short: "Create a new collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		// Create a new collection
		collectionName := args[0]
		newCollection, err := functions.NewCollection(collectionName)
		if err != nil {
			fmt.Printf("Error creating collection: %v\n", err)
			return
		}

		fmt.Printf("Created collection: %s\n", newCollection)
	},
}

// ioc collection rename
var collectionRenameCmd = &cobra.Command{
	Use:   "rename <collection_name> <new_collection_name>",
	Short: "Rename a collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}

		collectionName := args[0]
		newCollectionName := args[1]

		collection, err := functions.RenameCollection(collectionName, newCollectionName)
		if err != nil {
			fmt.Printf("Error renaming collection: %v\n", err)
			return
		}

		fmt.Printf("Renamed collection '%s' to '%s'\n", collectionName, collection.Name)
	},
}

// ioc collection delete
var collectionDeleteCmd = &cobra.Command{
	Use:   "delete <collection_name>",
	Short: "Delete a collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

// ioc collection import
var collectionImportCmd = &cobra.Command{
	Use:   "import <collection_name> <filepath>",
	Short: "Import a collection from a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}
	},
}

// ioc collection export
var collectionExportCmd = &cobra.Command{
	Use:   "export <filename>",
	Short: "Export a collection to a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}
	},
}

// ioc collection whitelist
var collectionWhitelistCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "Commands for collections with a whitelist",
}

// ioc collection whitelist add
var collectionWhitelistAddCmd = &cobra.Command{
	Use:   "add <whitelist_name>",
	Short: "Associate a whitelist to the selected collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

// ioc collection whitelist remove
var collectionWhitelistRemoveCmd = &cobra.Command{
	Use:   "remove <whitelist_name>",
	Short: "Disassociate a whitelist from the selected collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

func init() {
	collectionCmd.AddCommand(collectionListCmd)                     // ioc collection list
	collectionCmd.AddCommand(collectionCreateCmd)                   // ioc collection create
	collectionCmd.AddCommand(collectionDeleteCmd)                   // ioc collection delete
	collectionCmd.AddCommand(collectionImportCmd)                   // ioc collection import
	collectionCmd.AddCommand(collectionExportCmd)                   // ioc collection export
	collectionCmd.AddCommand(collectionRenameCmd)                   // ioc collection rename
	collectionWhitelistCmd.AddCommand(collectionWhitelistAddCmd)    // ioc collection whitelist add
	collectionWhitelistCmd.AddCommand(collectionWhitelistRemoveCmd) // ioc collection whitelist remove
	collectionCmd.AddCommand(collectionWhitelistCmd)                // ioc collection whitelist
	rootCmd.AddCommand(collectionCmd)                               // ioc collection
}
