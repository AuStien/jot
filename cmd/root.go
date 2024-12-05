package cmd

import (
	"github.com/austien/logbook/cmd/journal"
	"github.com/austien/logbook/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.PersistentFlags().String("home", "$HOME/.logbook", "home for logs (default is $HOME/.logbook)")
	RootCmd.PersistentFlags().String("editor", "vi", "which editor to use (default is vi")

	viper.SetEnvPrefix("LOGBOOK")
	viper.BindPFlag("home", RootCmd.LocalFlags().Lookup("home"))
	viper.BindEnv("home")
	viper.BindPFlag("editor", RootCmd.LocalFlags().Lookup("editor"))
	viper.BindEnv("editor", "EDITOR")

	RootCmd.AddCommand(todoCmd)
	RootCmd.AddCommand(binderCmd)
	RootCmd.AddCommand(journal.JournalCmd)

	config.Init()
}

var RootCmd = &cobra.Command{
	Use:   "logbook",
	Short: "Jot jot",
}

func Execute() error {
	return RootCmd.Execute()
}
