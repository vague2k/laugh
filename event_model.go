package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color(termANSIYellow.String())).
			Foreground(lipgloss.Color(termANSIBlack.String()))
	activePageStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIBrightWhite.String())).Render("•")
	summaryStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIYellow.String()))
	descriptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIBrightBlack.String()))
)

func NewEventListModel(events *[]Event) list.Model {
	// even though Event implements item, go treats 2 seperate slices differently,
	// even if both slices contain items that implement the same interface
	items := make([]list.Item, len(*events))
	for i, e := range *events {
		items[i] = e
	}

	l := list.New(items, NewEventDelegate(), width, height)
	l.Title = "Events"
	l.Paginator.ActiveDot = activePageStyle
	l.Styles.Title = titleStyle

	return l
}

// A "delegate" is bubbletea's (more specifically the bubbles list component) way of encapsulating functionality for list items.
//
// See the ItemDelegate interface and source code at https://github.com/charmbracelet/bubbles/list for more info.
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
	// the model's cursor has to be incremeneted/decremented manually because for some reason
	// m.CursorUp() and m.CursorDown() incremenet and decrement by 2 instead of 1?
	// this is probably a bug,
	// TODO: make a pr on this
	cursor := m.Cursor()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m.NewStatusMessage(fmt.Sprintf("%d", m.Index()))
		case "up", "k":
			if cursor > 0 {
				cursor--
			}
		case "down", "j":
			if cursor < len(m.Items())-1 {
				cursor++
			}
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

	if i, ok := item.(Event); ok {
		summary = i.Summary
		course = descriptionStyle.Render(i.Course)
		date = descriptionStyle.Render(i.GetFormattedStartDate())
	}

	cursor := summaryStyle.Render("│")
	if m.Index() == index {
		summary = summaryStyle.Render(summary)
		fmt.Fprintf(w, "%s %s\n%s %s\n%s %s", cursor, summary, cursor, course, cursor, date)
	} else {
		fmt.Fprintf(w, " %s\n %s\n %s", summary, course, date)
	}
}
