package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/kamontat/forgitgo/cmd/validation"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		err = validation.HasGit(dir)
		if err == nil {
			log.Fatalln(errors.New(".git already exist"))
		}

		_, err = git.PlainInit(dir, false)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
