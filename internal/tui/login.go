package tui

import (
	"os/exec"
	"runtime"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type loginModel struct {
	input  string
	width  int
	height int
	msg    string
}

func newLoginModel() loginModel {
	return loginModel{}
}

func (m loginModel) Init() tea.Cmd {
	return nil
}

func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			session := strings.TrimSpace(m.input)
			if session == "" {
				m.msg = ErrorStyle.Render("Session cannot be empty")
				return m, nil
			}
			return m, func() tea.Msg { return saveSessionMsg{session: session} }
		case "o":
			return m, openBrowserCmd
		case "esc":
			return m, func() tea.Msg { return goBackMsg{} }
		case "q":
			return m, tea.Quit
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		case "ctrl+u":
			m.input = ""
		default:
			if msg.Text != "" {
				m.input += msg.Text
			}
		}
	case tea.PasteMsg:
		m.input += msg.Content
	case loginSuccessMsg:
		m.msg = SuccessStyle.Render("Session saved successfully!")
		m.input = ""
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m loginModel) View() tea.View {
	lines := []string{
		TitleStyle.Render("Log In With LeetCode"),
		"",
		"1. Press 'o' to open LeetCode in your browser.",
		"2. Log in and copy your LEETCODE_SESSION cookie.",
		"3. Paste it below and press Enter.",
		"",
		"> " + m.input + "_",
		"",
	}
	if m.msg != "" {
		lines = append(lines, m.msg)
	}
	lines = append(lines, "", InfoStyle.Render("o open browser • esc back • q quit"))

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)
	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, BoxStyle.Render(content)))
	v.AltScreen = true
	return v
}

type saveSessionMsg struct{ session string }
type loginSuccessMsg struct{}

func openBrowserCmd() tea.Msg {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "start"
	default:
		cmd = "xdg-open"
	}
	args = append(args, "https://leetcode.com")
	_ = exec.Command(cmd, args...).Start()
	return nil
}
