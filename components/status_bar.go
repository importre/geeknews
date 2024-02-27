package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

func StatusBar(width int, a, b, c, d string) string {
	w := lipgloss.Width

	doc := strings.Builder{}
	statusKey := statusStyle.Render(a)
	encoding := encodingStyle.Render(b)
	fishCake := fishCakeStyle.Render(c)

	statusValWidth := width - w(statusKey) - w(encoding) - w(fishCake)
	statusVal := statusText.
		Width(statusValWidth).
		MaxHeight(1).
		Render(runewidth.Truncate(d, statusValWidth, "… "))

	bar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		statusKey,
		statusVal,
		encoding,
		fishCake,
	)

	doc.WriteString(statusBarStyle.Width(width).Render(bar))
	return doc.String()
}

var (
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	fishCakeStyle = statusNugget.Background(lipgloss.Color("#6124DF"))
)
