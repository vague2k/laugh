package main

import (
	"fmt"
	"math"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// tracks which part of the view is focused
type focusedState uint

const (
	listView focusedState = iota
	detailsView
)

type item struct{ title, desc string }

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// define pretty styles for TUI elements
var (
	// -3 is an offset to account for borders + help line at the bottom of the view
	width  = int(math.Round(float64(TermWidth())/2) - 3)
	height = TermHeight() - 3

	// TODO: use terminal colors instead of hardcoded values
	modelStyle = lipgloss.NewStyle().
			Width(width).
			Height(height).
			BorderStyle(lipgloss.HiddenBorder())

	focusedModelStyle = lipgloss.NewStyle().
				Width(width).
				Height(height).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("69"))

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(""))
)

type ParentModel struct {
	list    list.Model
	details list.Model
	focused focusedState
	Events  *[]Event
}

func newModel(events []*Event) ParentModel {
	// convert []*Event to []list.item
	var items []list.Item
	for _, event := range events {
		eventToItem := item{
			title: event.Summary,
			desc:  event.Course,
		}
		items = append(items, eventToItem)
	}
	// TODO: add more styling to list view elements, like item title, desc etc.
	// FIXME: FilterValue/filtering doesn't seem to do anything??? check docs
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Events"

	// HACK: for now, just to have something on the screen im using another list model
	// to show the right side of the view.
	//
	// TODO: the plan is to create a custom, but simple view where the event info can be seen in more detail
	detailsItems := []list.Item{
		item{title: "Item 1 details"},
		item{title: "Item 2 details"},
		item{title: "Item 3 details"},
	}
	d := list.New(detailsItems, list.NewDefaultDelegate(), width, height)
	d.Title = "Event details"

	model := ParentModel{
		list:    l,
		details: d,
		focused: listView,
		// Events: e,
	}

	return model
}

func (m ParentModel) Init() tea.Cmd {
	return nil
}

func (m ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.focused == listView {
				m.focused = detailsView
			} else {
				m.focused = listView
			}
		}

		switch m.focused {
		case listView:
			m.list, cmd = m.list.Update(msg)
		case detailsView:
			m.details, cmd = m.details.Update(msg)
		}
	}
	return m, tea.Batch(cmd)
}

func (m ParentModel) View() string {
	var s string
	if m.focused == listView {
		s += lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(fmt.Sprintf("%4s", m.list.View())), modelStyle.Render(m.details.View()))
	} else {
		s += lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(fmt.Sprintf("%4s", m.list.View())), focusedModelStyle.Render(m.details.View()))
	}
	s += helpStyle.Render("\ntab: focus next • q: exit")
	return s
}
