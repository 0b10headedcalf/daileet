package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/0b10headedcalf/daileet/internal/models"
	"github.com/charmbracelet/lipgloss"
)

type reviewModel struct {
	problem models.Problem
	width   int
	height  int
}

func newReviewModel(p models.Problem) reviewModel {
	return reviewModel{problem: p}
}

func (m reviewModel) Init() tea.Cmd {
	return nil
}

func (m reviewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "1":
			return m, func() tea.Msg { return gradeMsg{problem: m.problem, result: 1} }
		case "2":
			return m, func() tea.Msg { return gradeMsg{problem: m.problem, result: 2} }
		case "3":
			return m, func() tea.Msg { return gradeMsg{problem: m.problem, result: 3} }
		case "4":
			return m, func() tea.Msg { return gradeMsg{problem: m.problem, result: 4} }
		case "esc":
			return m, func() tea.Msg { return goBackMsg{} }
		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m reviewModel) View() tea.View {
	p := m.problem
	lines := []string{
		TitleStyle.Render("Review Problem"),
		"",
		fmt.Sprintf("Title:       %s", p.Title),
		fmt.Sprintf("Difficulty:  %s", p.Difficulty),
		fmt.Sprintf("Pattern:     %s", p.Pattern),
		fmt.Sprintf("URL:         %s", p.URL),
		"",
	}

	if p.LastReviewed != nil {
		lines = append(lines, fmt.Sprintf("Last reviewed: %s", p.LastReviewed.Format("2006-01-02")))
		lines = append(lines, fmt.Sprintf("Repetitions:   %d", p.Repetitions))
		lines = append(lines, fmt.Sprintf("Ease factor:   %.2f", p.EaseFactor))
		lines = append(lines, "")
	}

	lines = append(lines, PromptStyle.Render("How did it go?"))
	lines = append(lines,
		ErrorStyle.Render("1")+" Again  |  "+
			PromptStyle.Render("2")+" Hard  |  "+
			SuccessStyle.Render("3")+" Good  |  "+
			SuccessStyle.Render("4")+" Easy",
	)
	lines = append(lines, "")
	lines = append(lines, InfoStyle.Render("esc to cancel • q to quit"))

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)
	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, BoxStyle.Render(content)))
	v.AltScreen = true
	return v
}

type gradeMsg struct {
	problem models.Problem
	result  int
}
