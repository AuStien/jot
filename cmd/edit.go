package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the any file, creating directories if necessary",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		levels := strings.Split(args[0], "/")
		if len(levels) > 1 {
			path := []string{}
			path = append(path, j.HomeDir, "binder")
			path = append(path, levels[:len(levels)-1]...)
			if err := os.MkdirAll(filepath.Join(path...), 0755); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

		if err := j.Editor.OpenFile(filepath.Join(j.HomeDir, "binder", fmt.Sprintf("%s.md", args[0]))); err != nil {
			fmt.Fprintf(os.Stderr, "failed editing %s.md: %s\n", args[0], err.Error())
			os.Exit(1)
		}
	},
}
