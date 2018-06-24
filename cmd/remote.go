package cmd

import (
	"errors"

	"github.com/kamontat/forgitgo/client"
	"github.com/kamontat/forgitgo/client/validation"
	"github.com/kamontat/forgitgo/utils"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/config"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:     "remote",
	Aliases: []string{"r"},
	Short:   "Git cloud remote management",
	Run: func(cmd *cobra.Command, args []string) {
		runner := client.BeforeRun().SetRepository()

		err := validation.HasRemote(runner.Repo)
		if err == nil {
			utils.Logger().WithError(errors.New("Remote already initial")).GitError("Remote", "Already exist")
		}

		_, err = runner.Repo.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: args,
		})
		if err != nil {
			utils.Logger().WithError(err).GitError("Remote", "Cannot create")
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}
