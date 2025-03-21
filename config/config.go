package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/vague2k/laugh/utils"
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
// [LoadConfig] makes sure that the config file exists, along with the directory
// it lives in. The directory is created in $XDG_DATA_HOME or
// $HOME/.local/share, whichever comes first.
func LoadConfig(name string) (*Config, error) {
	var dataDir string
	// for testing
	switch true {
	case name == "":
		dir, err := utils.UserConfigHome()
		if err != nil {
			return nil, confErr(err)
		}
		dataDir = dir
	default:
		dataDir = name
	}

	appDir := filepath.Join(dataDir, "laugh")
	// does nothing if the dir already exists
	err := os.MkdirAll(appDir, 0o777)
	if err != nil {
		return nil, err
	}

	// ensure config files exists
	file := filepath.Join(appDir, "laugh.toml")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		file, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		conf := &Config{
			Url: "the url to your student canvas calendar should go here",
		}
		e := toml.NewEncoder(file)
		if err := e.Encode(conf); err != nil {
			return nil, err
		}

	} else if err != nil {
		return nil, err
	}
	// does nothing if the dir already exists

	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return nil, confErr(err)
	}

	conf.dir = appDir
	conf.filename = file

	return &conf, nil
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
