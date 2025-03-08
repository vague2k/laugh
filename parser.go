package main

import (
	"bufio"
	"os"
	"strings"
)

// Parse parses calendar events from an iCalendar (.ics) file as specified by
// RFC 7986.
//
// If the iCalendar file isn't detected to be a student canvas calendar, It's
// treated as an error.
func Parse(fileName string) (*[]Event, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var Events []Event
	var event *Event
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
			event = &Event{}

		case strings.Contains(s, "SUMMARY"):
			processingSummary = true
			summary = strings.Split(s, "SUMMARY:")[1]

		case strings.Contains(s, "DESCRIPTION"):
			processingSummary = false
			desc = strings.Split(s, "DESCRIPTION:")[1]

		case strings.Contains(s, "DTSTART"):
			event.StartDate = strings.Split(s, ":")[1]

		case strings.Contains(s, "END:VEVENT"):
			event.Summary = strings.Split(summary, "[")[0]
			event.Description = desc

			// parse course from the summary
			left := strings.IndexRune(summary, '[')
			right := strings.IndexRune(summary, ']')
			// offset left by 1 because that char "[" is included in the string
			parsed := summary[left+1 : right]
			course := strings.ReplaceAll(parsed, "\\,", ",")
			event.Course = course

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
