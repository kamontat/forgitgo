package cmd

import (
	"github.com/kamontat/forgitgo/client"
	"github.com/kamontat/forgitgo/utils"
	"github.com/spf13/cobra"
)

var all bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add file or folder into git system",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetWorktree()

		if all {
			runner.Add(".")
		} else {
			for _, elem := range args {
				runner.Add(elem)
			}
		}

		utils.Logger().Info("status", "\n"+runner.Status().String())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&all, "all", "a", false, "add every untrack files and folders")
}
