package editors

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/austien/logbook/debug"
)

// editors is a map of all supported editors.
var editors = map[string]Editor{
	"vi":   Vi{},
	"vim":  Vim{},
	"nvim": Neovim{},
	"nano": Nano{},
}

// Editor contains all the methods required to
// use a specfic editor.
//
// Note: The editor doesn't need to do exactly what
// the method implies. For instance, if an editor doesn't
// support opening a file with the cursor at the end, just
// open the file normally and manually navigate to the bottom.
type Editor interface {
	GetEditorExecutable() string
	OpenFile(file string) error
	OpenFileWithCursorAtEnd(file string) error
	OpenFileReadOnly(file string) error
}

func GetEditor(name string) (Editor, error) {
	for k, v := range editors {
		if name == k {
			return v, nil
		}
	}

	return nil, fmt.Errorf("unsupported editor %q (feel free to create a PR adding support at https://github.com/austien/logbook)", name)
}

// executeCmd is a helper function for editors to execute commands.
func executeCmd(editor Editor, args ...string) error {
	cmd := exec.Command(editor.GetEditorExecutable(), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return debug.WithFrame(err)
	}

	return nil
}
