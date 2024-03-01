/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
		flag, err := cmd.Flags().GetBool("differentmessage")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if flag {
			fmt.Println("This is a different message")
			return
		}

		fmt.Println("Hello World!")
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
	rootCmd.Flags().BoolP("differentmessage", "d", false, "Toggle a different message")
}
