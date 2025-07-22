package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chris-bateman/OpsTerm/ui"
)

func main() {
	// Create the initial model from the auth selector
	authModel := ui.NewAuthSelector()

	// Enable full-screen terminal control (no scroll, better rendering)
	p := tea.NewProgram(authModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running OpsTerm: %v\n", err)
		os.Exit(1)
	}
}