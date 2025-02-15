package main

import (
	"fmt"
	"strings"
	"time"
)

type Event struct {
	StartDate   string
	Summary     string
	Description string
	Course      string
	Done        bool
}

func (e *Event) GetFormattedStartDate() string {
	var t time.Time
	// format includes the hour
	if strings.Contains(e.StartDate, "Z") {
		stamp, err := time.Parse("20060102T150405Z", e.StartDate)
		if err != nil {
			panic(err) // FIXME: handle this correctly idk how yet
		}
		t = stamp
		// format with no hour included
	} else {
		stamp, err := time.Parse("20060102", e.StartDate)
		if err != nil {
			panic(err) // FIXME: handle this correctly idk how yet
		}
		t = stamp
	}
	month := t.Format("Jan")
	day := t.Day()
	year := t.Year()
	hour := t.Format("3:04 PM") // if format includes the date only, hour is set to 12:00 AM, which is not a bug for this use case

	return fmt.Sprintf("%s %d, %d. %s", month, day, year, hour)
}

// NOTE: for formatting description, take into account escape char like "\n" or "\\" and handle word wrapping.
func (e *Event) GetFormattedDescription() string {
	return ""
}
