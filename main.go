package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	parser, err := NewEventParser("spring_2025_cal.ics")
	if err != nil {
		log.Fatal(err)
	}

	Events := parser.Parse()

	// tui model
	model := newModel(Events)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
