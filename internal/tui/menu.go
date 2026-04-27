package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type menuChoice int

const (
	menuDue menuChoice = iota
	menuSolved
	menuEditor
	menuLogin
	menuImportSolved
	menuPresets
	menuClearData
)

var menuItems = []string{
	"Due Problems",
	"See Previously Solved Problems",
	"Manually Edit Problem List",
	"Log In With LeetCode",
	"Import My Solved Problems",
	"Presets",
	"Clear All User Data",
}

type menuModel struct {
	cursor   int
	width    int
	height   int
	loggedIn bool
	msg      string
}

func newMenuModel(loggedIn bool) menuModel {
	return menuModel{cursor: 0, loggedIn: loggedIn}
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
	case setAuthStatusMsg:
		m.loggedIn = msg.loggedIn
	case importSolvedResultMsg:
		if msg.err != nil {
			m.msg = ErrorStyle.Render(fmt.Sprintf("Import failed: %v", msg.err))
		} else {
			m.msg = SuccessStyle.Render(fmt.Sprintf("Imported %d solved problems (%d new, %d updated)", msg.total, msg.added, msg.updated))
		}
	case clearDataResultMsg:
		if msg.err != nil {
			m.msg = ErrorStyle.Render(fmt.Sprintf("Clear failed: %v", msg.err))
		} else {
			m.msg = SuccessStyle.Render("All user data cleared.")
			m.loggedIn = false
		}
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

	authIndicator := InfoStyle.Render("Not logged in")
	if m.loggedIn {
		authIndicator = SuccessStyle.Render("Logged in to LeetCode")
	}

	footer := InfoStyle.Render("j/k or ↑/↓ to navigate • Enter to select • q to quit")

	lines := []string{
		header,
		"",
		authIndicator,
		"",
		lipgloss.JoinVertical(lipgloss.Left, items...),
	}
	if m.msg != "" {
		lines = append(lines, "", m.msg)
	}
	lines = append(lines, "", footer)

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, BoxStyle.Render(content)))
	v.AltScreen = true
	return v
}

type menuSelectedMsg struct {
	choice menuChoice
}

type setAuthStatusMsg struct {
	loggedIn bool
}
