package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vague2k/laugh/parser"
)

type PagerModel struct {
	focused  string
	ready    bool
	viewport viewport.Model
}

func NewPagerModel() PagerModel {
	p := PagerModel{}
	p.viewport.HighPerformanceRendering = true
	return p
}

func (m PagerModel) Init() tea.Cmd {
	return nil
}

func (m PagerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case PagerMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Window.Width, msg.Window.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			if i, ok := msg.Focused.(parser.CalendarEvent); ok {
				m.viewport.SetContent(i.Description)
			}
			m.ready = true
		} else {
			if i, ok := msg.Focused.(parser.CalendarEvent); ok {
				m.viewport.SetContent(i.Description)
			}
			m.viewport.Width = msg.Window.Width
			m.viewport.Height = msg.Window.Height - verticalMarginHeight
		}
	case tea.KeyMsg:
		// Handle keyboard and mouse events in the viewport
		m.viewport, cmd = m.viewport.Update(msg)
	}

	return m, tea.Batch(cmd)
}

// TODO: add percent to show amount scrolled,
func (m PagerModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	// TODO: how to pad the viewport view without it bugging out??
	header := lipgloss.NewStyle().Padding(0, 1).Render(m.headerView())
	footer := lipgloss.NewStyle().Padding(0, 1).Render(m.footerView())
	return fmt.Sprintf("%s\n%s\n%s", header, m.viewport.View(), footer)
}

func (m PagerModel) headerView() string {
	s := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color(TermANSIBlack.String())).
		Background(lipgloss.Color(TermANSIYellow.String())).
		Render("Event Description")
	return fmt.Sprintf("%s\n", s)
}

func (m PagerModel) footerView() string {
	// TODO: Help should go here
	s := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color(TermANSIBrightBlack.String())).
		Render("Help here")
	return fmt.Sprintf("\n%s", s)
}
