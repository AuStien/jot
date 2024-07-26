package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	homeDir, ok := os.LookupEnv("LOGBOOK_HOME")
	if !ok {
		fmt.Fprintf(os.Stderr, "missing envvar LOGBOOK_HOME\n")
		os.Exit(1)
	}

	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		editor = "vim"
	}

	var isViewOnly bool
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "view":
			isViewOnly = true
		default:
			fmt.Fprintf(os.Stderr, "unknown command %s\n", os.Args[1])
			os.Exit(1)
		}
	}

	time := time.Now()

	dirPath := filepath.Join(homeDir, strconv.Itoa(time.Year()), fmt.Sprintf("%02d", time.Month()))

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create directory %q: %s\n", dirPath, err.Error())
		os.Exit(1)
	}

	filePath := filepath.Join(dirPath, fmt.Sprintf("%02d.md", time.Day()))
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "unknown error when stating %q: %s\n", filePath, err.Error())
			os.Exit(1)
		}

		file, err = os.Create(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create file %q: %s\n", filePath, err.Error())
			os.Exit(1)
		}

		_, err = file.Write([]byte(fmt.Sprintf("# %s %02d/%02d/%d\n", time.Weekday().String(), time.Day(), int(time.Month()), time.Year())))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to write header to file %q: %s\n", filePath, err.Error())
			file.Close()
			os.Exit(1)
		}
	}

	if !isViewOnly {
		if _, err := file.Write([]byte(fmt.Sprintf("\n## %02d:%02d\n\n", time.Hour(), time.Minute()))); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write subheader to file %q: %s\n", filePath, err.Error())
			file.Close()
			os.Exit(1)
		}
	}

	file.Close()

	cmd := exec.Command(editor, "+", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run command: %v\n", err)
		os.Exit(1)
	}
}
