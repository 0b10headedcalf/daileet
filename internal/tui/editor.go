package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/0b10headedcalf/daileet/internal/models"
	"github.com/charmbracelet/lipgloss"
)

type editorMode int

const (
	editorBrowse editorMode = iota
	editorAddInput
)

type editorModel struct {
	mode         editorMode
	problems     []models.Problem
	cursor       int
	scrollOffset int
	width        int
	height       int
	input        string
	msg          string
}

func newEditorModel() editorModel {
	return editorModel{mode: editorBrowse}
}

func (m editorModel) Init() tea.Cmd {
	return func() tea.Msg { return refreshEditorMsg{} }
}

func (m editorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refreshEditorMsg:
		return m, func() tea.Msg { return doRefreshEditorMsg{} }
	case doRefreshEditorMsg:
		return m, loadEditorProblemsCmd
	case loadedEditorProblemsMsg:
		m.problems = msg.problems
		m.cursor = 0
		m.scrollOffset = 0
		m.msg = ""
	case tea.PasteMsg:
		if m.mode == editorAddInput {
			m.input += msg.Content
		}
	case addProblemSuccessMsg:
		m.mode = editorBrowse
		m.input = ""
		m.msg = SuccessStyle.Render("Added: " + msg.title)
		return m, func() tea.Msg { return refreshEditorMsg{} }
	case tea.KeyPressMsg:
		if m.mode == editorAddInput {
			return m.handleAddInput(msg)
		}
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
		case "a":
			m.mode = editorAddInput
			m.input = ""
			m.msg = InfoStyle.Render("Enter LeetCode title slug (e.g. two-sum), then press Enter:")
		case "d":
			if len(m.problems) > 0 && m.cursor < len(m.problems) {
				p := m.problems[m.cursor]
				return m, func() tea.Msg { return deleteProblemMsg{id: p.ID, title: p.Title} }
			}
		case "esc":
			if m.mode == editorBrowse {
				return m, func() tea.Msg { return goBackMsg{} }
			}
			m.mode = editorBrowse
			m.input = ""
			m.msg = ""
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

func (m editorModel) handleAddInput(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		slug := strings.TrimSpace(m.input)
		if slug == "" {
			m.msg = ErrorStyle.Render("Title slug cannot be empty")
			return m, nil
		}
		return m, func() tea.Msg { return addProblemMsg{slug: slug} }
	case "esc":
		m.mode = editorBrowse
		m.input = ""
		m.msg = ""
	case "backspace":
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
		}
	case "ctrl+u":
		m.input = ""
	default:
		if msg.Text != "" && msg.Text != " " {
			m.input += msg.Text
		}
	}
	return m, nil
}

func (m editorModel) View() tea.View {
	var list []string
	start := m.scrollOffset
	end := start + m.maxVisible()
	if end > len(m.problems) {
		end = len(m.problems)
	}
	for i := start; i < end; i++ {
		p := m.problems[i]
		line := fmt.Sprintf("%s  %s", difficultyBadge(p.Difficulty), p.Title)
		if i == m.cursor {
			list = append(list, SelectedMenuItemStyle.Render("> "+line))
		} else {
			list = append(list, MenuItemStyle.Render("  "+line))
		}
	}
	if len(m.problems) == 0 {
		list = append(list, InfoStyle.Render("  No problems."))
	}

	var inputBox string
	if m.mode == editorAddInput {
		inputBox = BoxStyle.Render(m.msg + "\n" + "> " + m.input + "_")
	} else if m.msg != "" {
		inputBox = m.msg
	}

	var header string
	if len(m.problems) > m.maxVisible() {
		header = fmt.Sprintf("a add • d delete • ↑/↓/pgup/pgdown • %d/%d • esc back • q quit", m.cursor+1, len(m.problems))
	} else {
		header = "a add • d delete selected • esc back • q quit"
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		TitleStyle.Render("Problem Editor"),
		InfoStyle.Render(header),
		"",
		lipgloss.JoinVertical(lipgloss.Left, list...),
		"",
		inputBox,
	)

	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, BoxStyle.Render(content)))
	v.AltScreen = true
	return v
}

func (m editorModel) maxVisible() int {
	const overhead = 12 // borders, padding, title, header, empty lines, inputBox
	v := m.height - overhead
	if v < 1 {
		return 1
	}
	return v
}

func (m *editorModel) clampScroll() {
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

type refreshEditorMsg struct{}
type doRefreshEditorMsg struct{}
type loadedEditorProblemsMsg struct{ problems []models.Problem }
type deleteProblemMsg struct{ id int; title string }
type addProblemMsg struct{ slug string }
type addProblemSuccessMsg struct{ title string }

func loadEditorProblemsCmd() tea.Msg {
	return loadedEditorProblemsMsg{problems: nil}
}
