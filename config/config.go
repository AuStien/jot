package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/austien/jot/editors"
	"github.com/spf13/viper"
)

type Config struct {
	RootDir string
	Editor  editors.Editor
}

var config Config

func Init() {
	editor, err := editors.GetEditor(viper.GetString("editor"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed getting editor setup for %s: %s\n", viper.GetString("editor"), err.Error())
		os.Exit(1)
	}

	rootDir, err := filepath.Abs(viper.GetString("home"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed getting absolute path for %s: %s\n", viper.GetString("home"), err.Error())
		os.Exit(1)
	}

	config = Config{
		RootDir: rootDir,
		Editor:  editor,
	}
}

func Get() Config {
	return config
}
