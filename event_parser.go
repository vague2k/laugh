package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type EventParser struct {
	fileName string
	scanner  *bufio.Scanner
}

var ErrNotICSFile = errors.New("the file provided is not an .ics file")

func NewEventParser(fileName string) (*EventParser, error) {
	file, openPathErr := os.Open(fileName)
	if openPathErr != nil {
		return nil, openPathErr
	}

	if filepath.Ext(fileName) != ".ics" {
		return nil, ErrNotICSFile
	}

	// NOTE: preferrably add a check if it's a canvas calendar
	return &EventParser{
		fileName: fileName,
		scanner:  bufio.NewScanner(file),
	}, nil
}

// The name of the course can always be found inside the summary,
// enclosed by a "[" and "]"
func (p *EventParser) parseCourse(s string) string {
	left := strings.IndexRune(s, '[')
	right := strings.IndexRune(s, ']')
	// offset left by +1 because that char ("[") is included in the string
	parsed := s[left+1 : right]
	course := strings.ReplaceAll(parsed, "\\,", ",")
	return course
}

func (p *EventParser) Parse() []*Event {
	var Events []*Event
	var event *Event
	var summary string
	var desc string
	processingSummary := false

	for p.scanner.Scan() {
		s := p.scanner.Text()

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

		// NOTE: descriptions contain string chars like "\" or "\n". this has to be handled eventually to properly format in calendar view
		case strings.Contains(s, "DESCRIPTION"):
			processingSummary = false
			desc = strings.Split(s, "DESCRIPTION:")[1]

		case strings.Contains(s, "DTSTART"):
			event.StartDate = strings.Split(s, ":")[1]

		case strings.Contains(s, "END:VEVENT"):
			event.Summary = strings.Split(summary, "[")[0]
			event.Description = desc
			event.Course = p.parseCourse(summary)
			Events = append(Events, event)

			// cleanup
			event = nil
			desc = ""
			summary = ""
		}
	}

	if err := p.scanner.Err(); err != nil {
		panic(err)
	}

	return Events
}
