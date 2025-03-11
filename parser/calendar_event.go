package parser

type CalendarEvent struct {
	DueDate     string
	DueHour     string
	Summary     string
	Description string
	Course      string
	Done        bool
}

// this function is here to conform to the [list.Item] interface.
func (e CalendarEvent) FilterValue() string {
	return e.Summary
}
