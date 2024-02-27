package topic

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"importre.com/geeknews/components"
	"importre.com/geeknews/geeknews"
	"importre.com/geeknews/utils"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case geeknews.TopicRequest:
		return m, tea.Batch(m.setLoading(), m.fetch(msg))

	case Loading:
		m.loading = true
		return m, nil

	case FetchTopicResponse:
		m.topic = msg.topic
		m.render()
		m.viewport.GotoTop()
		m.loading = false
		return m, nil

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		h, v := screenStyle.GetFrameSize()
		m.viewport.Width = m.width - h
		m.viewport.Height = m.height - v - 1
		m.render()
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.clear()
			return m, nil

		case "v":
			if len(m.topic.Url) > 0 {
				utils.Open(m.topic.Url)
			}
			return m, nil

		case "ctrl+c":
			return m, tea.Quit

		default:
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}

	default:
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	if m.loading {
		return m.center("로딩중...")
	}

	return screenStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.viewport.View(),
			components.StatusBar(
				m.width-screenStyle.GetHorizontalFrameSize(),
				fmt.Sprintf("%dP | %v", m.topic.Info.Points, m.topic.Info.User),
				fmt.Sprintf("%d%%", int(m.viewport.ScrollPercent()*100)),
				fmt.Sprintf("%v", m.topic.Info.Timestamp),
				m.topic.Title,
			),
		),
	)
}
