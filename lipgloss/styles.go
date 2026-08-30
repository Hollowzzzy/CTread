package lipgloss

import "charm.land/lipgloss/v2"

var ERR = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Red)

var SUCCESS = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Green)
var INFO = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.White)
