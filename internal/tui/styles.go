package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	ColorPrimary   = lipgloss.Color("#A78BFA")
	ColorSecondary = lipgloss.Color("#34D399")
	ColorDanger    = lipgloss.Color("#F87171")
	ColorWarning   = lipgloss.Color("#FBBF24")
	ColorMuted     = lipgloss.Color("#6B7280")
	ColorBg        = lipgloss.Color("#1F2937")

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			MarginBottom(1)

	MenuItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2)

	SelectedMenuItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2).
			Foreground(ColorBg).
			Background(ColorPrimary).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorDanger)

	PromptStyle = lipgloss.NewStyle().
			Foreground(ColorWarning).
			Bold(true)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 2)
)
