package journal

import (
	"github.com/spf13/cobra"
)

func init() {
	JournalCmd.AddCommand(addCmd)
	JournalCmd.AddCommand(viewCmd)
}

var JournalCmd = &cobra.Command{
	Use:     "journal",
	Aliases: []string{"j"},
	Short:   "Handle the journal",
}
