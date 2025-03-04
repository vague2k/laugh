package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// EventListModel is a wrapper over the list component, where the list this
// model uses under the hood, uses an "EventDelegate" for custom functionality
// for list items.
type EventListModel struct {
	list list.Model
}

func NewEventListModel(events *[]Event) EventListModel {
	// even though Event implements item, go treats 2 seperate slices
	// differently even if both slices contain items that implement the same
	// interface
	items := make([]list.Item, len(*events))
	for i, e := range *events {
		items[i] = e
	}

	d := NewEventDelegate()
	l := list.New(items, d, width, height)
	l.Title = "Events"
	l.Paginator.ActiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.Color(termANSIBrightWhite.String())).
		Render("•")
	l.Paginator.InactiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.Color(termANSIBrightBlack.String())).
		Render("•")
	l.Styles.Title = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color(termANSIBrightYellow.String())).
		Foreground(lipgloss.Color(termANSIBlack.String()))

	return EventListModel{
		list: l,
	}
}

// A "delegate" is bubbletea's (more specifically the bubbles list component)
// way of encapsulating functionality for list items.
//
// See the ItemDelegate interface and source code at
// https://github.com/charmbracelet/bubbles/list for more info.
type EventDelegate struct {
	height, spacing int
}

func NewEventDelegate() EventDelegate {
	return EventDelegate{
		height:  3, // how many lines an item takes up in the view
		spacing: 1, // gap inbetween each item in the view
	}
}

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

func (d EventDelegate) Spacing() int {
	return d.spacing
}

func (d EventDelegate) Height() int {
	return d.height
}

func (d EventDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var summary, course, date string

	hoveredStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIBrightYellow.String()))
	descriptionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIBrightBlack.String()))

	if i, ok := item.(Event); ok {
		summary = i.Summary
		course = descriptionStyle.Render(i.Course)
		date = descriptionStyle.Render(i.GetFormattedStartDate())
	}

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
