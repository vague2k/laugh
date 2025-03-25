package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vague2k/laugh/config"
	"github.com/vague2k/laugh/database"
	"github.com/vague2k/laugh/models"
	"github.com/vague2k/laugh/parser"
)

type App struct {
	db   database.DBInterface
	conf *config.Config
	tui  *tea.Program
	err  error
}

func NewApp(conf *config.Config, db database.DBInterface) *App {
	return &App{
		db:   db,
		conf: conf,
	}
}

func (a *App) Run() {
	Events, err := parser.Parse("spring_2025_cal.ics")
	if err != nil {
		a.err = err
	}

	a.tui = tea.NewProgram(
		models.NewGlobalModel(Events),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion())

	if _, err := a.tui.Run(); err != nil {
		a.err = err
	}
}

func (a App) Err() error {
	if a.err != nil {
		return a.err
	}
	return nil
}
