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
	A long description for the CLI tool for Kanban board
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
