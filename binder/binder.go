package binder

import (
	"path/filepath"

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
