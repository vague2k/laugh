package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Event struct {
	StartDate   string
	Summary     string
	Description string
	Course      string
}

func main() {
	scanner, err := openICSFile("spring_2025_cal.ics")
	if err != nil {
		panic(err)
	}

	var Events []*Event
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

		// NOTE: descriptions contain string chars like "\" or "\n". this has to be handled eventually to properly format in calendar view
		case strings.Contains(s, "DESCRIPTION"):
			processingSummary = false
			desc = strings.Split(s, "DESCRIPTION:")[1]

		case strings.Contains(s, "DTSTART"):
			date := strings.Split(s, ":")[1]
			event.StartDate = parseDate(date)

		case strings.Contains(s, "END:VEVENT"):
			event.Summary = strings.Split(summary, "[")[0]
			event.Description = desc
			event.Course = parseCourse(summary)
			Events = append(Events, event)

		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, e := range Events {
		fmt.Printf("\nStart: %s\nCourse: %s\nSummary: %s\nDescription: %s\n",
			e.StartDate,
			// e.EndDate,
			e.Course,
			e.Summary,
			e.Description)
	}
}

// FIXME: check if file is canvas calendar ("Canvas" in "X-WR-CALNAME:")
func openICSFile(f string) (*bufio.Scanner, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, nil
}

func parseCourse(s string) string {
	left := strings.IndexRune(s, '[')
	right := strings.IndexRune(s, ']')
	return s[left+1 : right]
}

func parseDate(s string) string {
	var t time.Time
	// format includes the hour
	if strings.Contains(s, "Z") {
		stamp, err := time.Parse("20060102T150405Z", s)
		if err != nil {
			panic(err)
		}
		t = stamp
		// format with no hour included
	} else {
		stamp, err := time.Parse("20060102", s)
		if err != nil {
			panic(err)
		}
		t = stamp
	}
	month := t.Format("Jan")
	day := t.Day()
	year := t.Year()
	hour := t.Format("3:04 PM") // if format includes the date only, hour is set to 12:00 AM, which is not a bug for this use case

	result := fmt.Sprintf("%s %d, %d. %s", month, day, year, hour)
	return result
}
