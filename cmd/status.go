package cmd

import (
	"github.com/kamontat/forgitgo/client"
	"github.com/kamontat/forgitgo/utils"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Show git status of each file",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetRepository().SetWorktree()
		utils.Logger().Info("status", "\n"+runner.Status().String())
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
