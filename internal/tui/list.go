package tui

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/0b10headedcalf/daileet/internal/models"
	"github.com/charmbracelet/lipgloss"
)

type problemListKind int

const (
	listDue problemListKind = iota
	listSolved
)

type problemListModel struct {
	kind         problemListKind
	problems     []models.Problem
	cursor       int
	scrollOffset int
	width        int
	height       int
	err          string
}

func newProblemListModel(kind problemListKind) problemListModel {
	return problemListModel{kind: kind}
}

func (m problemListModel) Init() tea.Cmd {
	return func() tea.Msg {
		return refreshListMsg{kind: m.kind}
	}
}

func (m problemListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refreshListMsg:
		if msg.kind != m.kind {
			return m, nil
		}
		return m, func() tea.Msg { return doRefreshListMsg{kind: m.kind} }
	case doRefreshListMsg:
		if msg.kind != m.kind {
			return m, nil
		}
		return m, loadProblemsCmd(m.kind)
	case loadedProblemsMsg:
		if msg.kind != m.kind {
			return m, nil
		}
		m.problems = msg.problems
		m.cursor = 0
		m.scrollOffset = 0
		m.err = ""
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.problems)-1 {
				m.cursor++
			}
		case "pgup":
			visible := m.maxVisible()
			m.cursor -= visible
			if m.cursor < 0 {
				m.cursor = 0
			}
		case "pgdown":
			visible := m.maxVisible()
			m.cursor += visible
			if m.cursor >= len(m.problems) {
				m.cursor = len(m.problems) - 1
			}
		case "home", "g":
			m.cursor = 0
		case "end", "G":
			m.cursor = len(m.problems) - 1
		case "enter":
			if len(m.problems) > 0 && m.cursor < len(m.problems) {
				return m, func() tea.Msg {
					return reviewProblemMsg{problem: m.problems[m.cursor]}
				}
			}
		case "esc":
			return m, func() tea.Msg { return goBackMsg{} }
		case "q":
			return m, tea.Quit
		}
		m.clampScroll()
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m problemListModel) View() tea.View {
	var title string
	switch m.kind {
	case listDue:
		title = "Due Problems"
	case listSolved:
		title = "Previously Solved"
	}

	var list []string
	start := m.scrollOffset
	end := start + m.maxVisible()
	if end > len(m.problems) {
		end = len(m.problems)
	}
	for i := start; i < end; i++ {
		p := m.problems[i]
		line := fmt.Sprintf("%s  %s", difficultyBadge(p.Difficulty), p.Title)
		if m.kind == listSolved {
			line += fmt.Sprintf("  (rep: %d, ef: %.2f)", p.Repetitions, p.EaseFactor)
		} else if p.DueDate != nil {
			days := int(p.DueDate.Sub(time.Now()).Hours() / 24)
			if days < 0 {
				line += fmt.Sprintf("  overdue %dd", -days)
			} else {
				line += fmt.Sprintf("  in %dd", days)
			}
		} else {
			line += "  new"
		}
		if i == m.cursor {
			list = append(list, SelectedMenuItemStyle.Render("> "+line))
		} else {
			list = append(list, MenuItemStyle.Render("  "+line))
		}
	}

	if len(m.problems) == 0 {
		list = append(list, InfoStyle.Render("  No problems here."))
	}

	var footer string
	if len(m.problems) > m.maxVisible() {
		footer = fmt.Sprintf("↑/↓/pgup/pgdown • %d/%d • Enter to review • esc back • q quit", m.cursor+1, len(m.problems))
	} else {
		footer = "j/k or ↑/↓ to navigate • Enter to review • esc to go back • q to quit"
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		TitleStyle.Render(title),
		"",
		lipgloss.JoinVertical(lipgloss.Left, list...),
		"",
		InfoStyle.Render(footer),
	)

	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, BoxStyle.Render(content)))
	v.AltScreen = true
	return v
}

func (m problemListModel) maxVisible() int {
	const overhead = 10 // borders, padding, title, empty lines, footer
	v := m.height - overhead
	if v < 1 {
		return 1
	}
	return v
}

func (m *problemListModel) clampScroll() {
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor >= len(m.problems) {
		m.cursor = len(m.problems) - 1
		if m.cursor < 0 {
			m.cursor = 0
		}
	}
	visible := m.maxVisible()
	if m.cursor < m.scrollOffset {
		m.scrollOffset = m.cursor
	}
	if m.cursor >= m.scrollOffset+visible {
		m.scrollOffset = m.cursor - visible + 1
	}
}

func difficultyBadge(d models.Difficulty) string {
	switch d {
	case models.Easy:
		return SuccessStyle.Render("[E]")
	case models.Medium:
		return PromptStyle.Render("[M]")
	case models.Hard:
		return ErrorStyle.Render("[H]")
	default:
		return InfoStyle.Render("[?]")
	}
}

// Messages
type refreshListMsg struct{ kind problemListKind }
type doRefreshListMsg struct{ kind problemListKind }
type loadedProblemsMsg struct {
	kind     problemListKind
	problems []models.Problem
}
type reviewProblemMsg struct{ problem models.Problem }
type goBackMsg struct{}

func loadProblemsCmd(kind problemListKind) tea.Cmd {
	return func() tea.Msg {
		// Placeholder: the app router will intercept this and load via DB.
		// We return a nil msg so the app can inject the real data.
		return loadedProblemsMsg{kind: kind, problems: nil}
	}
}
