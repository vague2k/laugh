package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// The program's configuration options, which are loaded into by [LoadConfig]
//
// NOTE: fields that are private are meant for internal use, usually tests.
type Config struct {
	Url      string // the url where the .ics file will be requested from
	dir      string // the directory where the program's config file lives
	filename string // the program's config file's name
}

func (c Config) Dir() string {
	return c.dir
}

func (c Config) Filename() string {
	return c.filename
}

// Loads values from "laugh.toml" for use through out the program
//
// LoadConfig makes sure that the config file exists, along with the directory
// it lives in. The directory is created in $XDG_DATA_HOME or
// $HOME/.local/share, whichever comes first.
func LoadConfig(name string) (*Config, error) {
	var dataDir string
	// for testing
	switch true {
	case name == "":
		dir, err := userConfigHome()
		if err != nil {
			return nil, confErr(err)
		}
		dataDir = dir
	default:
		dataDir = name
	}

	// does nothing if the dir already exists
	appDir := filepath.Join(dataDir, "laugh")
	err := os.MkdirAll(appDir, 0o777)
	if err != nil {
		return nil, confErr(err)
	}

	file := filepath.Join(appDir, "laugh.toml")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		if err := createFile(file); err != nil {
			return nil, confErr(err)
		}
	} else if err != nil {
		return nil, confErr(err)
	}

	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return nil, confErr(err)
	}

	conf.dir = appDir
	conf.filename = file

	return &conf, nil
}

// creates a .toml with default values
func createFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	conf := &Config{
		Url: "the url to your student canvas calendar should go here",
	}
	e := toml.NewEncoder(file)
	if err := e.Encode(conf); err != nil {
		return err
	}
	return nil
}

// Gets the user's $XDG_CONFIG_HOME dir.
//
// NOTE: this is a stripped down and modified version of [os.UserConfigDir], the
// reason the ladder wasn't used was because on macOS, the default config
// directory is ".../Library/Application", this sucks, so we try to use
// $XDG_CONFIG_HOME or $HOME/.config instead, whichever one is found first
func userConfigHome() (string, error) {
	var dir string
	dir = os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		dir = os.Getenv("HOME")
		if dir == "" {
			return "", confErr("neither $XDG_CONFIG_HOME nor $HOME are defined")
		}
		dir += "/.config"
	} else if !filepath.IsAbs(dir) {
		return "", confErr("path in $XDG_CONFIG_HOME is relative")
	}

	return dir, nil
}

func confErr(msg any) error {
	switch msg.(type) {
	case string:
		return fmt.Errorf("Config: %v", msg)
	case error:
		return fmt.Errorf("Config: %v", msg)
	default:
		return nil
	}
}
