package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type presetChoice int

const (
	presetBlind75 presetChoice = iota
	presetGrind75
	presetCustomJSON
)

var presetItems = []string{
	"Blind 75",
	"Grind 75",
	"Custom JSON",
}

type presetsMode int

const (
	presetsMenu presetsMode = iota
	presetsInputPath
	presetsLoading
)

type presetsModel struct {
	mode   presetsMode
	cursor int
	input  string
	msg    string
	width  int
	height int
}

func newPresetsModel() presetsModel {
	return presetsModel{mode: presetsMenu, cursor: 0}
}

func (m presetsModel) Init() tea.Cmd {
	return nil
}

func (m presetsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if m.mode == presetsInputPath {
			return m.handleInputPath(msg)
		}
		if m.mode == presetsLoading {
			// Block all input while loading
			return m, nil
		}
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(presetItems)-1 {
				m.cursor++
			}
		case "enter":
			choice := presetChoice(m.cursor)
			if choice == presetCustomJSON {
				m.mode = presetsInputPath
				m.input = ""
				m.msg = InfoStyle.Render("Enter path to JSON file with problem slugs, then press Enter:")
				return m, nil
			}
			m.mode = presetsLoading
			m.msg = ""
			return m, func() tea.Msg { return loadPresetMsg{choice: choice} }
		case "esc":
			return m, func() tea.Msg { return goBackMsg{} }
		case "q":
			return m, tea.Quit
		}
	case presetLoadedMsg:
		m.mode = presetsMenu
		if msg.failed > 0 {
			m.msg = ErrorStyle.Render(fmt.Sprintf("Added %d, failed %d", msg.added, msg.failed))
			if len(msg.failedSlugs) > 0 {
				failedList := strings.Join(msg.failedSlugs, ", ")
				if len(failedList) > 60 {
					failedList = failedList[:57] + "..."
				}
				m.msg += "\n" + InfoStyle.Render("Failed: "+failedList)
			}
		} else {
			m.msg = SuccessStyle.Render(fmt.Sprintf("Added %d problems!", msg.added))
		}
	case tea.PasteMsg:
		if m.mode == presetsInputPath {
			m.input += msg.Content
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m presetsModel) handleInputPath(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		path := strings.TrimSpace(m.input)
		if path == "" {
			m.msg = ErrorStyle.Render("Path cannot be empty")
			return m, nil
		}
		m.mode = presetsLoading
		m.msg = ""
		return m, func() tea.Msg { return loadCustomPresetMsg{path: path} }
	case "esc":
		m.mode = presetsMenu
		m.input = ""
		m.msg = ""
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
	return m, nil
}

func (m presetsModel) View() tea.View {
	var content string

	if m.mode == presetsLoading {
		lines := []string{
			TitleStyle.Render("Presets"),
			"",
			PromptStyle.Render("Adding problems, please wait..."),
			"",
			InfoStyle.Render("This may take a few seconds"),
		}
		content = lipgloss.JoinVertical(lipgloss.Center, lines...)
	} else if m.mode == presetsInputPath {
		lines := []string{
			TitleStyle.Render("Custom Preset"),
			"",
			m.msg,
			"> " + m.input + "_",
			"",
			InfoStyle.Render("esc back • enter confirm • q quit"),
		}
		content = lipgloss.JoinVertical(lipgloss.Left, lines...)
	} else {
		items := make([]string, len(presetItems))
		for i, item := range presetItems {
			if i == m.cursor {
				items[i] = SelectedMenuItemStyle.Render("> " + item)
			} else {
				items[i] = MenuItemStyle.Render("  " + item)
			}
		}

		lines := []string{
			TitleStyle.Render("Presets"),
			"",
			lipgloss.JoinVertical(lipgloss.Left, items...),
		}
		if m.msg != "" {
			lines = append(lines, "", m.msg)
		}
		lines = append(lines, "", InfoStyle.Render("j/k or ↑/↓ to navigate • Enter to select • esc back • q quit"))
		content = lipgloss.JoinVertical(lipgloss.Center, lines...)
	}

	v := tea.NewView(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, BoxStyle.Render(content)))
	v.AltScreen = true
	return v
}

type loadPresetMsg struct {
	choice presetChoice
}

type loadCustomPresetMsg struct {
	path string
}

type presetLoadedMsg struct {
	added       int
	failed      int
	failedSlugs []string
}
