package parser

import (
	"os"
	"strings"

	"golang.org/x/term"
)

type CalendarEvent struct {
	DueDate     string
	DueHour     string
	Summary     string
	Description string
	Course      string
	Done        bool
}

// this function is here to conform to the [list.Item] interface.
func (e CalendarEvent) FilterValue() string {
	return e.Summary
}

func (e CalendarEvent) GetFormattedDescription() string {
	// HACK: this works for now, but width would eventually have to be based on
	// the window where the description is being shown in a TUI.
	termWidth, _, err := term.GetSize(int(os.Stderr.Fd()))
	if err != nil {
		panic(err)
	}
	wrappedDesc := e.wrapWords(e.Description, termWidth)

	escapeCharReplacer := strings.NewReplacer("\\n", "\n", "\\t", "\t", "\\,", ",")
	descWithReplacedEscChar := escapeCharReplacer.Replace(wrappedDesc)

	return descWithReplacedEscChar
}

// NOTE: this word wrap implementation MIGHT be cut, depending on if the
// bubbletea TUI view has word wrapping builtin
func (e CalendarEvent) wrapWords(s string, width int) string {
	words := strings.Fields(s)
	wrappedLen := 0
	var wrappedWords string

	for _, word := range words {
		// start new line
		if wrappedLen+len(word) > width {
			wrappedWords = strings.TrimSpace(wrappedWords) + "\n" + word
			wrappedLen = len(word) + 1
		} else {
			// otherwise append to current line
			if wrappedLen > 0 {
				wrappedWords += " "
			}
			wrappedWords += word
			wrappedLen += len(word) + 1
		}
	}

	return wrappedWords
}
