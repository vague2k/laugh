package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Parse parses calendar events from an iCalendar (.ics) file as specified by
// RFC 7986.
//
// If the iCalendar file isn't detected to be a student canvas calendar, It's
// treated as an error.
func Parse(fileName string) (*[]CalendarEvent, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var Events []CalendarEvent
	var event *CalendarEvent
	var summary string
	var desc string
	processingSummary := false

	for scanner.Scan() {
		s := scanner.Text()

		switch true {
		case strings.HasPrefix(s, " "):
			if processingSummary {
				summary += s[1:]
			} else {
				desc += s[1:]
			}

		case strings.Contains(s, "BEGIN:VEVENT"):
			event = &CalendarEvent{}

		case strings.Contains(s, "SUMMARY"):
			processingSummary = true
			summary = strings.Split(s, "SUMMARY:")[1]

		case strings.Contains(s, "DESCRIPTION"):
			processingSummary = false
			desc = strings.Split(s, "DESCRIPTION:")[1]

		case strings.Contains(s, "DTSTART"):
			timestamp := strings.Split(s, ":")[1]
			date, hour := parseDateAndHour(timestamp)
			event.DueDate = date
			event.DueHour = hour

		case strings.Contains(s, "END:VEVENT"):
			event.Summary = strings.Split(summary, "[")[0]
			event.Description = desc
			event.Course = parseCourse(summary)

			Events = append(Events, *event)

			// cleanup
			event = nil
			desc = ""
			summary = ""
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return &Events, nil
}

func parseCourse(s string) string {
	// parse course from the summary
	left := strings.IndexRune(s, '[')
	right := strings.IndexRune(s, ']')
	// offset left by 1 because that char "[" is included in the string
	parsed := s[left+1 : right]
	course := strings.ReplaceAll(parsed, "\\,", ",")
	return course
}

func parseDateAndHour(s string) (string, string) {
	var t time.Time

	// set the timezone location
	//
	// TODO: perhaps this could be configured in a config? I don't think anyone
	// else would use this besides me but just in case someone does idk
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return "Unknown Date", "Unknown Hour"
	}

	// format includes the hour
	if strings.Contains(s, "Z") {
		stamp, err := time.Parse("20060102T150405Z", s)
		if err != nil {
			return "Unknown Date", "Unknown Hour"
		}
		t = stamp.In(loc)
		// format with no hour included
	} else {
		stamp, err := time.Parse("20060102", s)
		if err != nil {
			return "Unknown Date", "Unknown Hour"
		}
		// No need to set loc for this timestamp since the hour being set to
		// midnight is intended behavior
		t = stamp
	}

	date := fmt.Sprintf("%s %d, %d", t.Format("Jan"), t.Day(), t.Year())
	hour := t.Format("3:04 PM")

	return date, hour
}
