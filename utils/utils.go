package utils

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/term"
)

// Gets the user's $XDG_CONFIG_HOME dir.
//
// NOTE: this is a stripped down and modified version of [os.UserConfigDir], the
// reason the ladder wasn't used was because on macOS, the default config
// directory is ".../Library/Application", this sucks, so we try to use
// $XDG_CONFIG_HOME or $HOME/.config instead, whichever one is found first
func UserConfigHome() (string, error) {
	var dir string
	dir = os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		dir = os.Getenv("HOME")
		if dir == "" {
			return "", errors.New("neither $XDG_CONFIG_HOME nor $HOME are defined")
		}
		dir += "/.config"
	} else if !filepath.IsAbs(dir) {
		return "", errors.New("path in $XDG_CONFIG_HOME is relative")
	}

	return dir, nil
}

// Gets the user's $XDG_DATA_HOME dir.
//
// Fallsback to the default data dir if the env var does not exist.
//
// NOTE: This implementation only supports unix right now
func UserDataHome() (string, error) {
	var dir string

	if dir = os.Getenv("XDG_DATA_HOME"); dir != "" {
		return dir, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir = home + "/.local/share"

	return dir, nil
}

func TermWidth() int {
	w, _ := getTerminalDimensions()
	return w
}

func TermHeight() int {
	_, h := getTerminalDimensions()
	return h
}

func getTerminalDimensions() (int, int) {
	w, h, err := term.GetSize(int(os.Stderr.Fd()))
	if err != nil {
		panic(err)
	}

	return w, h
}
