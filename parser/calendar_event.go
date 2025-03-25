package parser

import "strings"

type CalendarEvent struct {
	Id          int
	DueDate     string
	DueHour     string
	Summary     string
	Description string
	Course      string
	Done        bool
}

func (e CalendarEvent) WrapDescription(width int) string {
	wrapped := wrapWords(e.Description, width)
	rp := strings.NewReplacer("\\n", "\n", "\\t", "\t", "\\,", ",")
	desc := rp.Replace(wrapped)
	return desc
}

func wrapWords(s string, width int) string {
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

// this function is here to conform to the [list.Item] interface.
func (e CalendarEvent) FilterValue() string {
	return e.Summary
}
