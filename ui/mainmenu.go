package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuOption struct {
	Title       string
	Description string
	ActionID    string
}

type menuItem struct {
	option MenuOption
}

func (m menuItem) Title() string       { return m.option.Title }
func (m menuItem) Description() string { return m.option.Description }
func (m menuItem) FilterValue() string { return m.option.Title }

var menuOptions = []MenuOption{
	{"EC2 Instances", "View and manage EC2 instances", "ec2"},
	{"S3 Buckets", "List and inspect S3 buckets", "s3"},
	{"CloudWatch Logs", "Browse log groups and streams", "logs"},
	{"IAM Roles", "View and switch IAM roles", "iam"},
	{"Quit", "Exit OpsTerm", "quit"},
}

// Styling
var (
	menuBoxStyle = lipgloss.NewStyle().
		Margin(1, 2).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63"))

	menuTitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true).
		PaddingBottom(1)
)

type MainMenuModel struct {
	list list.Model
}

func NewMainMenu() tea.Model {
	items := make([]list.Item, len(menuOptions))
	for i, opt := range menuOptions {
		items[i] = menuItem{option: opt}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = menuTitleStyle.Render("OpsTerm: Main Menu")
	l.SetShowPagination(false)
	l.SetFilteringEnabled(false)

	return MainMenuModel{list: l}
}

func (m MainMenuModel) Init() tea.Cmd {
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := menuBoxStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			selected := m.list.SelectedItem().(menuItem).option
			switch selected.ActionID {
			case "ec2":
				// TODO: EC2 screen
			case "s3":
				// TODO: S3 screen
			case "logs":
				// TODO: CloudWatch screen
			case "iam":
				// TODO: IAM screen
			case "quit":
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MainMenuModel) View() string {
	return menuBoxStyle.Render(m.list.View())
}