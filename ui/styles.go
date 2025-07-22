// ui/styles.go
package ui

import "github.com/charmbracelet/lipgloss"

// Base layout
var BasePadding = lipgloss.NewStyle().
	Margin(1, 2)

// Title style
var TitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("212")).
	Bold(true).
	PaddingBottom(1).
	Align(lipgloss.Center)

// Highlighted list item
var SelectedItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("229")).
	Background(lipgloss.Color("57")).
	Bold(true)

// Description text
var DescriptionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("245"))

// Quit message / footer
var InfoStyle = lipgloss.NewStyle().
	MarginTop(1).
	Foreground(lipgloss.Color("240")).
	Italic(true)

// Error message
var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("9")).
	Bold(true)