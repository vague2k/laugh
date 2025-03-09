package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vague2k/laugh/models"
	"github.com/vague2k/laugh/parser"
)

func main() {
	// HACK: currently this works fine.
	// But to keep the calendar updated, this is my current plan.
	// 1. Download the calendar from the internet when program starts,
	// 2. If an .ics file DOESN'T exist, keep it, parse it, and put the parsed
	// content in a database for easy fetching later.
	// 3. If the a file DOES exist, compare byte content with the newly downloaded
	// .ics file,
	// 3.5 If there's a diff, parse it again and update info in the database
	// maybe there some way with this approach as to where the program doesn't
	// have to check the internet every single time on startup, perhaps it can
	// only do a check after a certain hour of the day a limited number of
	// times
	Events, err := parser.Parse("spring_2025_cal.ics")
	if err != nil {
		log.Fatal(err)
	}

	// tui model
	model := models.NewGlobalModel(Events)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
