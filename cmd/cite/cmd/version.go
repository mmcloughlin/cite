package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitSHA is intended to be populated with the git SHA of the repository at
// build time.
var gitSHA = "unknown"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print git revision",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(gitSHA)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
