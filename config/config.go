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

func Init() {
	editor, err := editors.GetEditor(viper.GetString("editor"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed getting editor setup for %s: %s\n", viper.GetString("editor"), err.Error())
		os.Exit(1)
	}

	home := os.ExpandEnv(viper.GetString("home"))
	if home == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		home = filepath.Join(userHomeDir, ".local", "share", "jot")
	} else {
		home, err = filepath.Abs(home)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed getting absolute path for %s: %s\n", viper.GetString("home"), err.Error())
			os.Exit(1)
		}
	}

	config = Config{
		HomeDir: home,
		Editor:  editor,
	}
}

func Get() Config {
	return config
}
