package cmd

import (
	"fmt"
	"os"

	"github.com/kamontat/forgitgo/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vp = viper.New()
var cfgFile string

// Username is commit username
var Username string

// Email commit email
var Email string

// CommitFormat is a format of commit message
var CommitFormat string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "forgitgo",
	Short:   "Formative git, wrote by pure go",
	Version: "1.0.0-rc.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.forgitgo.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Username, "name", "n", "", "commit name")
	rootCmd.PersistentFlags().StringVarP(&Email, "email", "e", "", "commit email")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		vp.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		vp.SetConfigType("yaml")
		vp.SetConfigName("config")
		vp.AddConfigPath(home + "/.forgitgo")
		vp.AddConfigPath("./.forgitgo")
	}

	vp.SetEnvPrefix("forgitgo")
	vp.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := vp.ReadInConfig(); err == nil {
		utils.Init(vp)

		utils.Logger().Debug("Config file", vp.ConfigFileUsed())

		Username = vp.GetString("user.name")
		Email = vp.GetString("user.email")
		CommitFormat = vp.GetString("commit.format")
	}
}
