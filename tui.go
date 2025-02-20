package main

import (
	"fmt"
)

// TODO: add mappings
// TODO: add config file
type Tui struct {
	Events []*Event
}

func NewTuiInstance(e []*Event) *Tui {
	return &Tui{
		Events: e,
	}
}

func (t *Tui) render() {
	for _, e := range t.Events {
		fmt.Printf("\nStart: %s\nCourse: %s\nSummary: %s\nDescription:\n%s\nDone: %v\n",
			e.GetFormattedStartDate(),
			e.Course,
			e.Summary,
			e.GetFormattedDescription(),
			e.Done,
		)
	}
}
