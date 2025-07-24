package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chris-bateman/OpsTerm/ui"
)

type appModel struct {
	current     tea.Model
	windowSize  *tea.WindowSizeMsg // cache the last known size
	switchingTo tea.Model          // temporary target for model switch
}

func (m appModel) Init() tea.Cmd {
	return m.current.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.windowSize = &msg // Cache it
		updated, cmd := m.current.Update(msg)
		m.current = updated
		return m, cmd

	case ui.SwitchToMainMenuMsg:
		newModel := ui.NewMainMenu()
		if m.windowSize != nil {
			// Send cached size to new model
			newModel, _ = newModel.Update(*m.windowSize)
		}
		m.current = newModel
		return m, nil
	}

	updated, cmd := m.current.Update(msg)
	m.current = updated
	return m, cmd
}

func (m appModel) View() string {
	return m.current.View()
}

func main() {
	root := appModel{current: ui.NewAuthSelector()}
	p := tea.NewProgram(root, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running OpsTerm: %v\n", err)
		os.Exit(1)
	}
}
