package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

type Event struct {
	StartDate   string
	Summary     string
	Description string
	Course      string
	Done        bool
}

func (e Event) FilterValue() string {
	return e.Summary
}

func (e Event) GetFormattedStartDate() string {
	var t time.Time
	// format includes the hour
	if strings.Contains(e.StartDate, "Z") {
		stamp, err := time.Parse("20060102T150405Z", e.StartDate)
		if err != nil {
			return "Unknown Time"
		}
		t = stamp
		// format with no hour included
	} else {
		stamp, err := time.Parse("20060102", e.StartDate)
		if err != nil {
			return "Unknown Time"
		}
		t = stamp
	}
	month := t.Format("Jan")
	day := t.Day()
	year := t.Year()
	hour := t.Format("3:04 PM") // if format includes the date only, hour is set to 12:00 AM, which is not a bug for this use case

	return fmt.Sprintf("%s %d, %d. %s", month, day, year, hour)
}

func (e Event) GetFormattedDescription() string {
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

// NOTE: this word wrap implementation MIGHT be cut, depending on if the bubbletea TUI view has word wrapping builtin
func (e Event) wrapWords(s string, width int) string {
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
