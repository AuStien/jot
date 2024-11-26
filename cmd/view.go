package cmd

import (
	"fmt"
	"os"

	"github.com/austien/logbook/journal"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the last entries in read-only",
	Run: func(cmd *cobra.Command, args []string) {
		j := journal.New(rootDir, editor)

		if _, err := j.ConcatLastMonth(); err != nil {
			fmt.Fprintf(os.Stderr, "getting last months files: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
