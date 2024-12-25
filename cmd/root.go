package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

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

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintln(os.Stderr, "failed to read build info")
		os.Exit(1)
	}

	revision := ""
	time := ""
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			revision = setting.Value
		case "vcs.time":
			time = setting.Value
		}
	}

	RootCmd.SetVersionTemplate(fmt.Sprintf("%s (%s)\n", revision, time))
}

var RootCmd = &cobra.Command{
	Use:     "jot",
	Short:   "Jot jot",
	Version: "devel",
}

func Execute() error {
	return RootCmd.Execute()
}
