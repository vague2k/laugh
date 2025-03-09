package models

import "fmt"

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

func (c termANSIColor) String() string {
	return fmt.Sprintf("%d", c)
}
