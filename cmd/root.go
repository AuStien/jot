package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/austien/logbook/book"
	"github.com/austien/logbook/editors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var b book.Book

func init() {
	rootCmd.PersistentFlags().String("home", "$HOME/.logbook", "home for logs (default is $HOME/.logbook)")
	rootCmd.PersistentFlags().String("editor", "vi", "which editor to use (default is vi")

	viper.SetEnvPrefix("LOGBOOK")
	viper.BindPFlag("home", rootCmd.LocalFlags().Lookup("home"))
	viper.BindEnv("home")
	viper.BindPFlag("editor", rootCmd.LocalFlags().Lookup("editor"))
	viper.BindEnv("editor", "EDITOR")

	rootCmd.AddCommand(todoCmd)
	rootCmd.AddCommand(viewCmd)
}

var rootCmd = &cobra.Command{
	Use:   "log",
	Short: "Jot jot",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		editor, err := editors.GetEditor(viper.GetString("editor"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed getting editor setup for %s: %s\n", viper.GetString("editor"), err.Error())
			os.Exit(1)
		}

		absoluteHomeDir, err := filepath.Abs(viper.GetString("home"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed getting absolute path for %s: %s\n", viper.GetString("home"), err.Error())
			os.Exit(1)
		}

		b = book.Book{
			HomeDir: absoluteHomeDir,
			Editor:  editor,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		if err := b.UpsertDayFile(now); err != nil {
			fmt.Fprintf(os.Stderr, "upserDayFile: %s\n", err.Error())
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
