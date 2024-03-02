/*
Copyright © 2024 Andrés Court andres@andrescourt.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/alcb1310/kanban/cmd/tui"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kanban",
	Short: "A CLI tool for Kanban board",
	Long: `
Manage efficiently your TODO items using a kanban board

ADDING A NEW ITEM:
	press n to add a new item, then write the title of the item and press enter

CHANGE THE STATUS OF AN ITEM:
	press space to change the status

EXIT:
	press ctrl + c to exit the kanban board

MOVE BETWEEN ITEMS IN A LIST:
	use the up and down arrow keys or j and k

MOVE BETWEEN LISTS:
	use the left and right arrow keys or h and l
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		tui.App()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Set different flags here
	// rootCmd.Flags().BoolP("differentmessage", "d", false, "Toggle a different message")
}
