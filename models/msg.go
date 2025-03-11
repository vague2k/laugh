package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type PagerMsg struct {
	Window  tea.WindowSizeMsg
	Focused list.Item
}
