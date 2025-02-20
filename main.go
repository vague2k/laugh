package main

func main() {
	parser, err := NewEventParser("spring_2025_cal.ics")
	if err != nil {
		panic(err)
	}

	Events := parser.Parse()

	tui := NewTuiInstance(Events)

	// NOTE: i know the current iteration just prints,
	// but i'd assume at some point this would also return an error?
	tui.render()
}
