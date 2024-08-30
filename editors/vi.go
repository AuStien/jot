package editors

// Vim editor, for more information see https://www.cs.colostate.edu/helpdocs/vi.html.
type Vi struct{}

var _ Editor = Vi{}

func (vi Vi) GetEditorExecutable() string {
	return "vi"
}

func (vi Vi) OpenFile(file string) error {
	return executeCmd(vi, file)
}

func (vi Vi) OpenFileWithCursorAtEnd(file string) error {
	return executeCmd(vi, "+", file)
}
func (vi Vi) OpenFileReadOnly(file string) error {
	return executeCmd(vi, "-R", file)
}
