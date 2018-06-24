package cmd

import (
	"github.com/kamontat/forgitgo/client"
	"github.com/kamontat/forgitgo/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add file or folder into git system",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetWorktree()

		for _, elem := range args {
			runner.Add(elem)
		}

		utils.Logger().Info("status", "\n"+runner.Status().String())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
