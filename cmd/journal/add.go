package journal

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/austien/jot/config"
	"github.com/austien/jot/journal"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry to the journal",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		cfg := config.Get()
		j := journal.New(cfg)

		if err := j.CreateEntry(now); err != nil {
			fmt.Fprintf(os.Stderr, "upserDayFile: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
