package book

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/austien/logbook/editors"
)

type Book struct {
	HomeDir string
	Editor  editors.Editor
}

// UpsertDayFile makes sure the file "2024/08/30.md"
// exists.
//
// If it was created the following header will be added:
//
//	# Friday 30/08/2024
//
// It then appends the following sub header:
//
//	## 14:35
//
// Lastly it opens the file for editing.
func (b Book) UpsertDayFile(at time.Time) error {
	year := fmt.Sprintf("%d", at.Year())
	month := fmt.Sprintf("%02d", at.Month())
	day := fmt.Sprintf("%02d", at.Day())

	rootDir := b.getJournalRootPath()

	if err := os.MkdirAll(filepath.Join(rootDir, year, month), 0755); err != nil {
		return err
	}

	filePath := filepath.Join(rootDir, year, month, fmt.Sprintf("%s.md", day))
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}

		file, err = os.Create(filePath)
		if err != nil {
			return fmt.Errorf("create file %s: %w", filePath, err)
		}

		_, err = file.Write([]byte(fmt.Sprintf("# %s %s/%s/%s\n", at.Weekday().String(), day, month, year)))
		if err != nil {
			file.Close()
			return fmt.Errorf("writing header to file %s: %w", filePath, err)
		}
	}
	if _, err := file.Write([]byte(fmt.Sprintf("\n## %02d:%02d\n\n", at.Hour(), at.Minute()))); err != nil {
		file.Close()
		return fmt.Errorf("writing subheader to file %s: %w", file.Name(), err)
	}

	if err := b.Editor.OpenFileWithCursorAtEnd(filePath); err != nil {
		return fmt.Errorf("editing file %s: %w", filePath, err)
	}

	return nil
}

// ConcatLastMonth concats all the files of the last month (with
// entries) in a temporary file and opens said file.
func (b Book) ConcatLastMonth() (string, error) {
	now := time.Now()
	rootDir := b.getJournalRootPath()

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return "", err
	}

	yearDir := ""
	year := ""
	// Only check the last 100 years
	for i := now.Year(); i >= now.Year()-100; i-- {
		year = fmt.Sprintf("%d", i)
		for _, entry := range entries {
			if entry.Name() == year {
				if !entry.IsDir() {
					return "", fmt.Errorf("%s/%s is not a dir", rootDir, entry)
				}

				yearDir = filepath.Join(rootDir, entry.Name())
				break
			}
		}
		if yearDir != "" {
			break
		}
	}

	if yearDir == "" {
		return "", fmt.Errorf("directory for year %d not found", now.Year())
	}

	entries, err = os.ReadDir(yearDir)
	if err != nil {
		return "", err
	}

	monthDir := ""
	month := ""
	for i := now.Month(); i > 0; i-- {
		month = fmt.Sprintf("%02d", i)
		for _, entry := range entries {
			if entry.Name() == month {
				if !entry.IsDir() {
					return "", fmt.Errorf("%s/%s is not a dir\n", yearDir, entry)
				}

				monthDir = filepath.Join(yearDir, entry.Name())
				break
			}
		}
		if monthDir != "" {
			break
		}
	}

	if monthDir == "" {
		return "", fmt.Errorf("directory for month %02d not found", now.Month())
	}

	tmpFile, err := os.CreateTemp("", fmt.Sprintf("logbook-%s-%s-*.md", year, month))
	if err != nil {
		return "", fmt.Errorf("createTemp: %w", err)
	}
	defer tmpFile.Close()

	entries, err = os.ReadDir(monthDir)
	if err != nil {
		return "", err
	}

	slices.SortFunc(entries, func(a, b fs.DirEntry) int {
		return strings.Compare(a.Name(), b.Name())
	})

	for _, entry := range entries {
		file, err := os.Open(filepath.Join(monthDir, entry.Name()))
		if err != nil {
			return "", err
		}
		defer file.Close()

		if _, err := tmpFile.ReadFrom(file); err != nil {
			return "", err
		}

		if _, err := tmpFile.WriteString(fmt.Sprintf("\n---\n\n")); err != nil {
			return "", err
		}
	}

	if err := b.Editor.OpenFileReadOnly(tmpFile.Name()); err != nil {
		return "", fmt.Errorf("openFileReadOnly %s: %w", tmpFile.Name(), err)
	}

	return tmpFile.Name(), nil
}

func (b Book) getJournalRootPath() string {
	return filepath.Join(b.HomeDir, "journal")
}
