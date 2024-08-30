package editors

// Neovim editor, for more information see https://neovim.io/.
type Neovim struct{}

var _ Editor = Neovim{}

func (nvim Neovim) GetEditorExecutable() string {
	return "nvim"
}

func (nvim Neovim) OpenFile(file string) error {
	return executeCmd(nvim, file)
}

func (nvim Neovim) OpenFileWithCursorAtEnd(file string) error {
	return executeCmd(nvim, "+", file)
}
func (nvim Neovim) OpenFileReadOnly(file string) error {
	return executeCmd(nvim, "-R", file)
}
