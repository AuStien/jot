package main

import (
	"fmt"
	"os"

	"github.com/austien/logbook/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "bruh: %s\n", err.Error())
		os.Exit(1)
	}
}
