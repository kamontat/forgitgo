package cmd

import (
	"github.com/kamontat/forgitgo/client"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Commit file or folder with format message",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetRepository().SetWorktree()

		runner.Commit(Username, Email, args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
