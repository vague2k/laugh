package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type PagerMsg struct {
	Window  tea.WindowSizeMsg
	Key     tea.KeyMsg
	Focused list.Item
}

func SendPagerMsg(width int, height int, focused list.Item) PagerMsg {
	msg := PagerMsg{}
	msg.Window.Width = width
	msg.Window.Height = height
	msg.Focused = focused
	return msg
}
