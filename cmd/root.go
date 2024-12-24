package cmd

import (
	"fmt"
	"os"

	"github.com/austien/jot/cmd/journal"
	"github.com/austien/jot/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.PersistentFlags().String("home", "$XDG_DATA_HOME", "home for notes (if $XDG_DATA_HOME isn't set, uses $HOME/.local/share/jot)")
	RootCmd.PersistentFlags().String("editor", "vi", "which editor to use")

	viper.SetEnvPrefix("JOT")
	viper.BindPFlag("home", RootCmd.LocalFlags().Lookup("home"))
	viper.BindEnv("home")
	viper.BindPFlag("editor", RootCmd.LocalFlags().Lookup("editor"))
	viper.BindEnv("editor", "EDITOR")

	RootCmd.AddCommand(todoCmd)
	RootCmd.AddCommand(binderCmd)
	RootCmd.AddCommand(journal.JournalCmd)

	if err := config.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "init: %s\n", err.Error())
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "jot",
	Short: "Jot jot",
}

func Execute() error {
	return RootCmd.Execute()
}
