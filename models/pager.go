package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vague2k/laugh/parser"
)

type PagerModel struct {
	ready    bool
	viewport viewport.Model
}

func NewPagerModel() PagerModel {
	p := PagerModel{}
	return p
}

func (m PagerModel) Init() tea.Cmd {
	return nil
}

// FIXME: event description not being shown in it's entirety in the viewport.
// I already checked the parser just in case it was an issue there, and it seems
// to not be
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
				m.viewport.SetContent(i.WrapDescription(m.viewport.Width))
			}
			m.ready = true
		} else {
			if i, ok := msg.Focused.(parser.CalendarEvent); ok {
				m.viewport.SetContent(i.WrapDescription(m.viewport.Width))
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

// TODO: how to pad the viewport itself without it bugging out??
func (m PagerModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	header := lipgloss.NewStyle().Padding(0, 1).Render(m.headerView())
	return fmt.Sprintf("%s\n%s\n%s", header, m.viewport.View(), m.footerView())
}

func (m PagerModel) headerView() string {
	s := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(TermANSIBlack.Color()).
		Background(TermANSIYellow.Color()).
		Render("Event Description")
	return fmt.Sprintf("%s\n", s)
}

// TODO: add help for mappings, similiar to the event list
func (m PagerModel) footerView() string {
	help := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(TermANSIBrightBlack.Color()).
		Render("Help here")

	scrolled := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(TermANSIBrightBlack.Color()).
		Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))

	maxLen := lipgloss.Width(help) + lipgloss.Width(scrolled)
	gap := strings.Repeat(" ", max(0, m.viewport.Width-maxLen))
	return lipgloss.JoinHorizontal(lipgloss.Center, help, gap, scrolled)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
