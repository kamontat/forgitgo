package cmd

import (
	"github.com/kamontat/forgitgo/client"
	"github.com/kamontat/forgitgo/utils"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run after create new project",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetRepository().SetWorktree()
		utils.Logger().Info("status", "\n"+runner.Status().String())
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
