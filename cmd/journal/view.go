package journal

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/austien/logbook/config"
	"github.com/austien/logbook/journal"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the last entries in read-only",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()
		j := journal.New(cfg)

		if _, err := j.ConcatLastMonth(); err != nil {
			fmt.Fprintf(os.Stderr, "getting last months files: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
