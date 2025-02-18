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

func (e *Event) GetFormattedDescription() string {
	escapeCharReplacer := strings.NewReplacer("\\n", "\n", "\\t", "\t", "\\,", ",")
	descWithReplacedEscChar := escapeCharReplacer.Replace(e.Description)

	// FIXME: handle word wrapping
	return descWithReplacedEscChar
}
