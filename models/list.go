package models

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vague2k/laugh/parser"
)

// NewEventListModel wraps over the list component, where the list this
// model uses under the hood, uses an [EventDelegate] for custom functionality
// for list items.
func NewEventListModel(events *[]parser.CalendarEvent) list.Model {
	// even though [CalendarEvent] implements [list.Item], go treats 2 seperate
	// slices differently even if both slices contain items that implement the
	// same interface
	items := make([]list.Item, len(*events))
	for i, e := range *events {
		items[i] = e
	}

	d := NewEventDelegate()
	l := list.New(items, d, width, height)
	l.Title = "Events"
	l.Paginator.ActiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.Color(TermANSIBrightWhite.String())).
		Render("•")
	l.Paginator.InactiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.Color(TermANSIBrightBlack.String())).
		Render("•")
	l.Styles.Title = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color(TermANSIBrightYellow.String())).
		Foreground(lipgloss.Color(TermANSIBlack.String()))

	return l
}

// A "delegate" is bubbletea's (more specifically the bubbles list component)
// way of encapsulating functionality for list items.
//
// See the [list.ItemDelegate] interface and source code at
// https://github.com/charmbracelet/bubbles/list for more info.
type EventDelegate struct {
	height, spacing int
}

func NewEventDelegate() EventDelegate {
	return EventDelegate{
		height:  3, // The height of the [list.Item] in the view.
		spacing: 1, // The space inbetween each [list.Item] in the view.
	}
}

// Update
// TODO: implement filtering
func (d EventDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m.NewStatusMessage(fmt.Sprintf("%d", m.Index()))
		}
	}
	return nil
}

// The space inbetween each [list.Item] in the view.
func (d EventDelegate) Spacing() int {
	return d.spacing
}

// The height of the [list.Item] in the view.
func (d EventDelegate) Height() int {
	return d.height
}

// Render handles how each [list.Item] is shown in the list view.
func (d EventDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var summary, course, date string

	hoveredStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(TermANSIBrightYellow.String()))
	descriptionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(TermANSIBrightBlack.String()))

	if i, ok := item.(parser.CalendarEvent); ok {
		summary = i.Summary
		course = descriptionStyle.Render(i.Course)
		date = descriptionStyle.Render(i.DueDate)
	}

	// Set styles differently if it's the currently focused item in the view.
	if m.Index() == index {
		cursor := hoveredStyle.Render("│")
		summary = hoveredStyle.Render(summary)
		fmt.Fprintf(w, "%s %s\n%s %s\n%s %s",
			cursor, summary,
			cursor, course,
			cursor, date)
	} else {
		fmt.Fprintf(w, " %s\n %s\n %s",
			summary,
			course,
			date)
	}
}
