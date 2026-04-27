package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type splashModel struct {
	width  int
	height int
}

func newSplashModel() splashModel {
	return splashModel{}
}

func (m splashModel) Init() tea.Cmd {
	return nil
}

func (m splashModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m, func() tea.Msg { return splashDoneMsg{} }
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m splashModel) View() tea.View {
	logo := `
██████╗  █████╗ ██╗██╗     ███████╗███████╗████████╗
██╔══██╗██╔══██╗██║██║     ██╔════╝██╔════╝╚══██╔══╝
██║  ██║███████║██║██║     █████╗  █████╗     ██║   
██║  ██║██╔══██║██║██║     ██╔══╝  ██╔══╝     ██║   
██████╔╝██║  ██║██║███████╗███████╗███████╗   ██║   
╚═════╝ ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝╚══════╝   ╚═╝   
`
	subtitle := "Spaced Repetition for LeetCode"
	hint := "Press any key to continue..."

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		TitleStyle.Render(logo),
		SubtitleStyle.Render(subtitle),
		"",
		InfoStyle.Render(hint),
	)

	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content))
	v.AltScreen = true
	return v
}

type splashDoneMsg struct{}
