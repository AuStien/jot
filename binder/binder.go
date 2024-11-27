package binder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/austien/logbook/editors"
)

const dirKey = "binder"

type binder struct {
	HomeDir string
	Editor  editors.Editor
}

func New(rootDir string, editor editors.Editor) binder {
	return binder{
		HomeDir: filepath.Join(rootDir, dirKey),
		Editor:  editor,
	}
}

func (b binder) AutoCompleteTargets(toComplete string) ([]string, error) {
	targets := []string{}
	levels := strings.Split(toComplete, string(os.PathSeparator))

	if len(levels) == 1 {
		entries, err := os.ReadDir(b.HomeDir)
		if err != nil {
			return nil, err
		}

		for _, ent := range entries {
			if strings.HasPrefix(ent.Name(), toComplete) {
				name := ent.Name()
				if ent.IsDir() {
					name += string(os.PathSeparator)
				}
				targets = append(targets, name)
			}
		}
	} else {
		leaf := levels[len(levels)-1]
		path := []string{}
		path = append(path, b.HomeDir)
		path = append(path, levels[:len(levels)-1]...)

		entries, err := os.ReadDir(filepath.Join(path...))
		if err != nil {
			return nil, err
		}

		for _, ent := range entries {
			if !ent.IsDir() && ent.Name() == leaf {
				return nil, nil
			}

			if strings.HasPrefix(ent.Name(), leaf) {
				name := ent.Name()
				if ent.IsDir() {
					name += string(os.PathSeparator)
				} else {
					name = strings.TrimPrefix(name, leaf)
				}
				targets = append(targets, fmt.Sprintf("%s%s", toComplete, name))
			}
		}
	}

	return targets, nil
}
