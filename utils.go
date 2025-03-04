package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func WriteStringf(builder *strings.Builder, s string, a ...any) {
	fmt.Fprintf(builder, s, a...)
}

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
