package ui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chris-bateman/OpsTerm/aws"
)

type AuthOption struct {
	Title       string
	Description string
	AuthType    int
}

type listItem struct {
	option AuthOption
}

func (li listItem) Title() string       { return li.option.Title }
func (li listItem) Description() string { return li.option.Description }
func (li listItem) FilterValue() string { return li.option.Title }

var options = []AuthOption{
	{"Use default AWS profile", "Uses the default profile in ~/.aws/config", 0},
	{"Use named AWS profile", "Prompt for a profile name", 1},
	{"Assume IAM role", "Prompt for role ARN and base profile", 2},
	{"Use environment variables", "Reads AWS creds from env vars", 3},
	{"Exit", "Quit the application", 4},
}

// 💅 Lipgloss styles
var (
	boxStyle = lipgloss.NewStyle().
			Margin(1, 2).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")) // blue-ish

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")). // pink
			Bold(true).
			PaddingBottom(1)
)

type authModel struct {
	list list.Model
}

func NewAuthSelector() tea.Model {
	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = listItem{option: opt}
	}

	delegate := list.NewDefaultDelegate()

	l := list.New(items, delegate, 0, 0)
	l.Title = titleStyle.Render("Select AWS Auth Method")
	l.SetShowPagination(false)
	l.SetFilteringEnabled(false)

	return authModel{list: l}
}

func (m authModel) Init() tea.Cmd {
	return nil
}

func (m authModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := boxStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			selected := m.list.SelectedItem().(listItem).option
			switch selected.AuthType {
			case 0:
				// Load AWS config using default profile
				cfg, err := aws.LoadAWSConfig(
					context.TODO(),
					aws.AuthInput{
						Method: aws.DefaultProfile,
						Region: "ap-southeast-2", // Default region
					},
				)
				if err != nil {
					m.list.Title = titleStyle.Render("Failed to load AWS config")
					return m, nil
				}
				fmt.Println("Loaded AWS config for region:", cfg.Region)

				// Send message to transition to main menu
				return m, func() tea.Msg { return SwitchToMainMenuMsg{} }

			case 1:
				// TODO: Named profile input
			case 2:
				// TODO: Assume role input
			case 3:
				// TODO: Use env vars
			case 4:
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m authModel) View() string {
	return boxStyle.Render(m.list.View())
}