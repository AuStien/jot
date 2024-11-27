package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/austien/logbook/binder"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the any file, creating directories if necessary",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b := binder.New(rootDir, editor)

		levels := strings.Split(args[0], "/")
		if len(levels) > 1 {
			path := []string{b.HomeDir}
			path = append(path, levels[:len(levels)-1]...)
			if err := os.MkdirAll(filepath.Join(path...), 0755); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

		// Make sure ".md" is added if no extension is specified.
		// If it is specify, make sure it isn't doubled up.
		ext := filepath.Ext(args[0])
		if ext == "" {
			ext = ".md"
		} else {
			ext = ""
		}

		if err := editor.OpenFile(filepath.Join(b.HomeDir, fmt.Sprintf("%s%s", args[0], ext))); err != nil {
			fmt.Fprintf(os.Stderr, "failed editing %s.md: %s\n", args[0], err.Error())
			os.Exit(1)
		}
	},
}
