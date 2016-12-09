package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command.
var RootCmd = &cobra.Command{
	Use:   "cite",
	Short: "Reference snippets in your godoc",
}

// Execute runs RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
