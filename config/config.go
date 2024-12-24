package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/austien/jot/editors"
	"github.com/spf13/viper"
)

type Config struct {
	HomeDir string
	Editor  editors.Editor
}

var config Config

func Init() error {
	editor, err := editors.GetEditor(viper.GetString("editor"))
	if err != nil {
		return fmt.Errorf("editor setup for %s: %w", viper.GetString("editor"), err)
	}

	home := os.ExpandEnv(viper.GetString("home"))
	if home == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("absolute path for %s: %w", viper.GetString("home"), err)
		}

		home = filepath.Join(userHomeDir, ".local", "share", "jot")
	} else {
		home, err = filepath.Abs(home)
		if err != nil {
			return fmt.Errorf("absolute path for %s: %w", viper.GetString("home"), err)
		}
	}

	config = Config{
		HomeDir: home,
		Editor:  editor,
	}

	return nil
}

func Get() Config {
	return config
}
