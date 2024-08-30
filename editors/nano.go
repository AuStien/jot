package editors

// Nano editor, for more information see https://www.nano-editor.org/.
type Nano struct{}

var _ Editor = Nano{}

func (nano Nano) GetEditorExecutable() string {
	return "nano"
}

func (nano Nano) OpenFile(file string) error {
	return executeCmd(nano, file)
}

func (nano Nano) OpenFileWithCursorAtEnd(file string) error {
	return executeCmd(nano, "+-1", file)
}
func (nano Nano) OpenFileReadOnly(file string) error {
	return executeCmd(nano, "--view", file)
}
