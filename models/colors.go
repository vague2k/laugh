package models

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type termANSIColor uint

const (
	TermANSIBlack termANSIColor = iota
	TermANSIRed
	TermANSIGreen
	TermANSIYellow
	TermANSIBlue
	TermANSIMagenta
	TermANSICyan
	TermANSIWhite

	TermANSIBrightBlack
	TermANSIBrightRed
	TermANSIBrightGreen
	TermANSIBrightYellow
	TermANSIBrightBlue
	TermANSIBrightMagenta
	TermANSIBrightCyan
	TermANSIBrightWhite
)

func (c termANSIColor) Color() lipgloss.Color {
	return lipgloss.Color(fmt.Sprintf("%d", c))
}
