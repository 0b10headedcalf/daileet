package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type menuChoice int

const (
	menuDue menuChoice = iota
	menuSolved
	menuEditor
	menuLogin
	menuPresets
)

var menuItems = []string{
	"Due Problems",
	"See Previously Solved Problems",
	"Manually Edit Problem List",
	"Log In With LeetCode",
	"Presets",
}

type menuModel struct {
	cursor int
	width  int
	height int
}

func newMenuModel() menuModel {
	return menuModel{cursor: 0}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(menuItems)-1 {
				m.cursor++
			}
		case "enter":
			return m, func() tea.Msg { return menuSelectedMsg{choice: menuChoice(m.cursor)} }
		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m menuModel) View() tea.View {
	items := make([]string, len(menuItems))
	for i, item := range menuItems {
		if i == m.cursor {
			items[i] = SelectedMenuItemStyle.Render("> " + item)
		} else {
			items[i] = MenuItemStyle.Render("  " + item)
		}
	}

	header := TitleStyle.Render("Daileet")
	footer := InfoStyle.Render("j/k or ↑/↓ to navigate • Enter to select • q to quit")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		"",
		lipgloss.JoinVertical(lipgloss.Left, items...),
		"",
		footer,
	)

	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content))
	v.AltScreen = true
	return v
}

type menuSelectedMsg struct {
	choice menuChoice
}
