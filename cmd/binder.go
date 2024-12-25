package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/austien/jot/binder"
	"github.com/austien/jot/config"
	"github.com/spf13/cobra"
)

var binderCmd = &cobra.Command{
	Use:     "binder",
	Aliases: []string{"b"},
	Short:   "Edit the any file in the binder, creating directories if necessary",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		b := binder.New(cfg)

		path := ""
		if len(args) > 0 {
			levels := strings.Split(args[0], string(os.PathSeparator))
			if len(levels) > 1 {
				path := []string{b.HomeDir}
				path = append(path, levels[:len(levels)-1]...)
				if err := os.MkdirAll(filepath.Join(path...), 0o755); err != nil {
					fmt.Fprintf(os.Stderr, "mkdir: %s\n", err.Error())
					os.Exit(1)
				}
			}
			path = args[0]
		}

		if err := cfg.Editor.OpenFile(filepath.Join(b.HomeDir, path)); err != nil {
			fmt.Fprintf(os.Stderr, "failed editing %s: %s\n", path, err.Error())
			os.Exit(1)
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		cfg := config.Get()

		b := binder.New(cfg)

		targets, err := b.AutoCompleteTargets(toComplete)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return nil, cobra.ShellCompDirectiveError
		}

		return targets, cobra.ShellCompDirectiveNoSpace
	},
}
