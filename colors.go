package main

import "fmt"

type termANSIColor uint

const (
	termANSIBlack termANSIColor = iota
	termANSIRed
	termANSIGreen
	termANSIYellow
	termANSIBlue
	termANSIMagenta
	termANSICyan
	termANSIWhite

	termANSIBrightBlack
	termANSIBrightRed
	termANSIBrightGreen
	termANSIBrightYellow
	termANSIBrightBlue
	termANSIBrightMagenta
	termANSIBrightCyan
	termANSIBrightWhite
)

func (c termANSIColor) String() string {
	return fmt.Sprintf("%d", c)
}
