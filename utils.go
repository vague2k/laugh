package main

import (
	"os"

	"golang.org/x/term"
)

func getTerminalDimensions() (int, int) {
	w, h, err := term.GetSize(int(os.Stderr.Fd()))
	if err != nil {
		panic(err)
	}

	return w, h
}

func TermWidth() int {
	w, _ := getTerminalDimensions()
	return w
}

func TermHeight() int {
	_, h := getTerminalDimensions()
	return h
}
