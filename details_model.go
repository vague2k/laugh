package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// the DetailsModel only serves to show an event's details in a pretty way, and
// nothing more
type DetailsModel struct {
	focused list.Item
	styles  DetailsStyles
}

func NewDetailsModel() tea.Model {
	s := DefaultDetailsStyles()
	return DetailsModel{
		styles: s,
	}
}

func (m DetailsModel) Init() tea.Cmd {
	return nil
}

func (m DetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case list.Item:
		m.focused = msg
	}
	return m, nil
}

func (m DetailsModel) View() string {
	var summary, course, date string

	if i, ok := m.focused.(CalendarEvent); ok {
		summary = i.Summary
		course = i.Course
		date = i.GetFormattedStartDate()
	}

	s := &strings.Builder{}
	fmt.Fprintf(s, "%s\n\n", m.styles.Title.Render("Event Details"))
	fmt.Fprintf(s, "\n%s %s\n",
		m.styles.DueDateLabel.Render("Due Date:"),
		m.styles.DueDate.Render(date))
	fmt.Fprintf(s, "\n%s %s\n",
		m.styles.SummaryLabel.Render("Classwork:"),
		m.styles.Summary.Render(summary))
	fmt.Fprintf(s, "\n%s %s\n",
		m.styles.CourseLabel.Render("For Course:"),
		m.styles.Course.Render(course))

	return m.styles.Model.Render(s.String())
}

// DetailsStyles contains style definitions for the details component. these
// values are generated by DefaultStyles.
type DetailsStyles struct {
	Title,
	Model,
	DueDate,
	DueDateLabel,
	Summary,
	SummaryLabel,
	Course,
	CourseLabel lipgloss.Style
}

// DefaultDetailsStyles generates styles for DetailsModel
func DefaultDetailsStyles() (s DetailsStyles) {
	s.Title = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color(termANSIBrightYellow.String())).
		Foreground(lipgloss.Color(termANSIBlack.String()))

	s.Model = lipgloss.NewStyle().Padding(0, 2)

	s.DueDate = lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIRed.String()))
	s.DueDateLabel = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color(termANSIRed.String())).
		Foreground(lipgloss.Color(termANSIWhite.String()))

	s.Summary = lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIWhite.String()))
	s.SummaryLabel = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color(termANSIBlack.String())).
		Foreground(lipgloss.Color(termANSIBrightWhite.String()))

	s.Course = lipgloss.NewStyle().Foreground(lipgloss.Color(termANSIWhite.String()))
	s.CourseLabel = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color(termANSIBlack.String())).
		Foreground(lipgloss.Color(termANSIBrightWhite.String()))

	return s
}
