package topic

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"importre.com/geeknews/geeknews"
)

func (m Model) setLoading() tea.Cmd {
	return func() tea.Msg {
		return Loading{}
	}
}

func (m Model) fetch(request geeknews.TopicRequest) tea.Cmd {
	return func() tea.Msg {
		response, _ := geeknews.FetchTopic(request)
		return FetchTopicResponse{
			topic: response.Topic,
		}
	}
}

func (m *Model) clear() {
	m.topic = geeknews.Topic{}
}

func (m *Model) render() {
	if !m.HasContent() {
		return
	}

	body := make(chan string)
	comments := make(chan string)

	go func() { body <- strings.TrimSpace(m.topic.Markdown(m.viewport.Width)) }()
	go func() { comments <- m.commentsView(m.viewport.Width) }()

	content := strings.Join([]string{<-body, <-comments}, "\n\n")
	m.viewport.SetContent(content)
}

func (m Model) commentsView(width int) string {
	commentsTitle := fmt.Sprintf("Comments [%d개]", len(m.topic.Comments))
	spaces := strings.Builder{}
	for i := 0; i < width-len(commentsTitle); i++ {
		spaces.WriteString(" ")
	}

	b := strings.Builder{}
	b.WriteString("\n")
	b.WriteString(commentsHeader(fmt.Sprintf("%s%s", commentsTitle, spaces.String())))
	b.WriteString("\n")

	for _, c := range m.topic.Comments {
		name := c.Info.User

		if name == m.topic.Info.User {
			name = nameStyle("* " + name)
		}

		reArrow := "  ↪  "
		arrowPrefix := ""
		for depth := 0; depth < max(int(c.Depth)-1, 0); depth++ {
			arrowPrefix += "    "
		}

		comment := fmt.Sprintf(
			"%s [%s]\n%s",
			name,
			c.Info.Timestamp,
			strings.TrimSpace(c.Markdown(width-len(arrowPrefix+reArrow))),
		)

		if c.Depth > 0 {
			b.WriteString(
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					arrowPrefix+reArrow,
					commentStyle.Render(comment),
				),
			)
		} else {
			b.WriteString(
				commentStyle.Render(comment),
			)
		}

		b.WriteString("\n")
	}

	return commentsStyle.Render(b.String())
}

func (m Model) center(s string) string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
}

var (
	screenStyle = lipgloss.NewStyle().
			Margin(2, 8)

	subtle         = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	commentsHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			Render

	commentsStyle = lipgloss.NewStyle()

	commentStyle = lipgloss.NewStyle().
			BorderBottom(true).
			BorderForeground(subtle).
			BorderStyle(lipgloss.NormalBorder())

	nameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false).
			Render
)
