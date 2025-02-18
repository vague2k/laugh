package main

import "fmt"

func main() {
	parser := NewEventParser("spring_2025_cal.ics")

	Events := parser.Parse()

	if err := parser.Err(); err != nil {
		panic(err)
	}

	for _, e := range Events {
		fmt.Printf("\nStart: %s\nCourse: %s\nSummary: %s\nDescription:\n%s\nDone: %v\n",
			e.GetFormattedStartDate(),
			e.Course,
			e.Summary,
			e.GetFormattedDescription(),
			e.Done,
		)
	}
}
