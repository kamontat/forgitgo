package cmd

import (
	"path"

	"github.com/kamontat/forgitgo/client"
	"github.com/spf13/cobra"
)

var addAll = false

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Commit file or folder with format message",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetRepository().SetWorktree()

		// TODO: check is user call 'git add' new file.

		runner.
			SetCommitDatabase(path.Join(path.Dir(vp.ConfigFileUsed()), "commit.yaml")).
			Commit(Username, Email, addAll, CommitFormat, args...)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().BoolVarP(&addAll, "add", "a", false, "Add -a to 'git commit'")
}
