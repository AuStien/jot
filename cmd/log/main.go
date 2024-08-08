package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

var homeDirFlag = flag.String("home", "", "Set home directory. Takes precedence over envvar")
var editorFlag = flag.String("editor", "", "Set editor. Takes precedence over envvar")

func main() {
	flag.Parse()

	homeDir := ""
	if homeDirFlag != nil && *homeDirFlag != "" {
		homeDir = *homeDirFlag
	}

	if homeDir == "" {
		homeDirEnv, ok := os.LookupEnv("LOGBOOK_HOME")
		if ok {
			homeDir = homeDirEnv
		}
	}
	if homeDir == "" {
		fmt.Fprintf(os.Stderr, "missing home dir. Set with LOGBOOK_HOME or --home\n")
		os.Exit(1)
	}

	editor := ""
	if editorFlag != nil && *editorFlag != "" {
		editor = *editorFlag
	}

	if editor == "" {
		editorEnv, ok := os.LookupEnv("EDITOR")
		if ok {
			editor = editorEnv
		} else {
			editor = "vim"
		}
	}

	var isViewOnly bool
	var isTodo bool
	if len(flag.Args()) > 0 {
		switch flag.Args()[0] {
		case "view":
			isViewOnly = true
		case "todo":
			isTodo = true
		default:
			fmt.Fprintf(os.Stderr, "unknown command %s\n", os.Args[0])
			os.Exit(1)
		}
	}

	if isTodo {
		cmd := exec.Command(editor, filepath.Join(homeDir, "TODO.md"))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to run command: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	now := time.Now()

	if isViewOnly {
		entries, err := os.ReadDir(homeDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readdir %q: %v\n", homeDir, err)
			os.Exit(1)
		}

		yearDir := ""
		// Only check the last 100 years
		for i := now.Year(); i >= now.Year()-100; i-- {
			year := fmt.Sprintf("%d", i)
			for _, entry := range entries {
				if entry.Name() == year {
					if !entry.IsDir() {
						fmt.Fprintf(os.Stderr, "\"%s/%s\" not a dir\n", homeDir, entry)
						os.Exit(1)
					}

					yearDir = filepath.Join(homeDir, entry.Name())
					goto month
				}
			}
		}

	month:
		if yearDir == "" {
			fmt.Fprintf(os.Stderr, "directory for year %d not found\n", now.Year())
			os.Exit(1)
		}

		entries, err = os.ReadDir(yearDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readdir %q: %v\n", yearDir, err)
			os.Exit(1)
		}

		monthDir := ""
		for i := now.Month(); i > 0; i-- {
			month := fmt.Sprintf("%02d", i)
			for _, entry := range entries {
				if entry.Name() == month {
					if !entry.IsDir() {
						fmt.Fprintf(os.Stderr, "\"%s/%s\" not a dir\n", yearDir, entry)
						os.Exit(1)
					}

					monthDir = filepath.Join(yearDir, entry.Name())
					goto day
				}
			}
		}

	day:
		if monthDir == "" {
			fmt.Fprintf(os.Stderr, "directory for month %02d not found\n", now.Month())
			os.Exit(1)
		}

		entries, err = os.ReadDir(monthDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readdir %q: %v\n", monthDir, err)
			os.Exit(1)
		}

		dayFile := ""
		for i := now.Day(); i > 0; i-- {
			day := fmt.Sprintf("%02d.md", i)
			for _, entry := range entries {
				if entry.Name() == day {
					dayFile = filepath.Join(monthDir, entry.Name())
					goto open
				}
			}
		}

	open:
		if dayFile == "" {
			fmt.Fprintf(os.Stderr, "file for day %02d not found\n", now.Day())
			os.Exit(1)
		}
		cmd := exec.Command(editor, "+", dayFile)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to run command: %v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	dirPath := filepath.Join(homeDir, strconv.Itoa(now.Year()), fmt.Sprintf("%02d", now.Month()))

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create directory %q: %s\n", dirPath, err.Error())
		os.Exit(1)
	}

	filePath := filepath.Join(dirPath, fmt.Sprintf("%02d.md", now.Day()))
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

		_, err = file.Write([]byte(fmt.Sprintf("# %s %02d/%02d/%d\n", now.Weekday().String(), now.Day(), int(now.Month()), now.Year())))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to write header to file %q: %s\n", filePath, err.Error())
			file.Close()
			os.Exit(1)
		}
	}

	if _, err := file.Write([]byte(fmt.Sprintf("\n## %02d:%02d\n\n", now.Hour(), now.Minute()))); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write subheader to file %q: %s\n", filePath, err.Error())
		file.Close()
		os.Exit(1)
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
