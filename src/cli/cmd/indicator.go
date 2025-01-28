package cmd

import (
	"fmt"

	"github.com/egrzeszczak/ioc/src/cli/functions"
	"github.com/spf13/cobra"
)

// ioc list
var listCmd = &cobra.Command{
	Use:   "list <collection_name>",
	Short: "List all indicators in a collection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		collectionName := args[0]

		// Get all indicators in collection
		indicators, err := functions.GetIndicators(collectionName)
		if err != nil {
			fmt.Printf("Error getting indicators: %v\n", err)
		}

		for _, indicator := range indicators {
			fmt.Printf("%v\n", indicator)
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

		collectionName := args[0]
		indicator := args[1]

		// Add indicator to collection
		newIndicator, err := functions.NewIndicator(collectionName, indicator)
		if err != nil {
			fmt.Printf("Error adding indicator: %v\n", err)
			return
		}

		fmt.Printf("Added indicator: %s\n", newIndicator)
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
	rootCmd.AddCommand(listCmd)    // ioc list
	rootCmd.AddCommand(addCmd)     // ioc add
	rootCmd.AddCommand(bulkAddCmd) // ioc bulk-add
	rootCmd.AddCommand(searchCmd)  // ioc search
}
