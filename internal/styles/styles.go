package styles

import (
	"charm.land/lipgloss/v2"
)

var (
	white = lipgloss.Color("#FFFFFF")
	alpha = lipgloss.Color("#7f23db")
)

var (
	// Message box style
	MainMsgBoxStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground((white)).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground((alpha)).
			MarginLeft(2)

	// Main logo style
	MainAsciiLogoStyle = lipgloss.NewStyle().
				Bold(false).
				Foreground((white))
)
