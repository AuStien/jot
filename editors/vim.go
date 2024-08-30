package editors

// Vim editor, for more information see https://www.vim.org/.
type Vim struct{}

var _ Editor = Vim{}

func (vim Vim) GetEditorExecutable() string {
	return "vim"
}

func (vim Vim) OpenFile(file string) error {
	return executeCmd(vim, file)
}

func (vim Vim) OpenFileWithCursorAtEnd(file string) error {
	return executeCmd(vim, "+", file)
}
func (vim Vim) OpenFileReadOnly(file string) error {
	return executeCmd(vim, "-R", file)
}
