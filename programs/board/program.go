package board

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"importre.com/geeknews/components"
	"importre.com/geeknews/geeknews"
)

func (m Model) Init(key string, page uint) tea.Cmd {
	return func() tea.Msg {
		response, _ := geeknews.FetchBoard(geeknews.BoardRequest{Key: key, Page: page})
		return FetchBoardResponse{
			page:  page,
			board: response.Board,
		}
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.setTableLayout()
		h, v := screenStyle.GetFrameSize()
		m.boards.SetSize(msg.Width-h, msg.Height-v)
		m.topic, cmd = m.topic.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if m.topic.HasContent() {
			m.topic, cmd = m.topic.Update(msg)
			return m, cmd
		}

		if m.visibleBoards() {
			return m.update(msg)
		}

		switch msg.String() {
		case "r":
			return m, m.Init(m.board.Key, m.page)

		case "d":
			if m.board.IsLastPage {
				return m, nil
			}
			return m, m.Init(m.board.Key, m.page+1)

		case "u":
			if m.page <= 1 {
				return m, nil
			}
			return m, m.Init(m.board.Key, m.page-1)

		case "o":
			items := make([]list.Item, len(boardItems))
			for i, boardItem := range boardItems {
				items[i] = boardItem
			}
			m.boards.SetItems(items)
			return m, nil

		case "enter":
			if len(m.board.Topics) > 0 {
				t := m.board.Topics[m.table.Cursor()]
				topicMsg := geeknews.TopicRequest{Id: t.Id}
				m.topic, cmd = m.topic.Update(topicMsg)
			}
			return m, cmd

		case "q", "ctrl+c":
			return m, tea.Quit

		default:
			m.table, cmd = m.table.Update(msg)
			cmds = append(cmds, cmd)
			m.topic, cmd = m.topic.Update(msg)
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)
		}

	case FetchBoardResponse:
		rows := make([]table.Row, len(msg.board.Topics))
		for i, topic := range msg.board.Topics {
			rows[i] = table.Row{
				fmt.Sprintf("%4d", topic.Info.Points),
				topic.Title,
				fmt.Sprintf("%4d", topic.Info.NumComments),
				topic.Info.User,
				topic.Info.Timestamp,
			}
		}

		if len(rows) <= m.table.Cursor() {
			m.table.SetCursor(len(rows) - 1)
		}

		m.page = msg.page
		m.table.SetRows(rows)
		m.table, cmd = m.table.Update(msg)
		m.board = msg.board
		m.loading = false
		m.setTableLayout()
		m.boards.SetItems([]list.Item{})
		return m, cmd

	default:
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)
		m.topic, cmd = m.topic.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}
}

func (m Model) View() string {
	if m.loading {
		return m.center("로딩중...")
	}

	if m.topic.HasContent() {
		return m.topic.View()
	}

	if len(m.boards.Items()) > 0 {
		return screenStyle.Render(
			m.boards.View(),
		)
	}

	lastPage := ""
	if m.board.IsLastPage {
		lastPage = " (마지막)"
	}

	preview := ""
	cursor := m.table.Cursor()
	if len(m.board.Topics) > cursor && cursor >= 0 {
		preview = m.board.Topics[cursor].Preview
	}

	return screenStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.table.View(),
			components.StatusBar(
				m.width-screenStyle.GetHorizontalFrameSize(),
				m.title,
				fmt.Sprintf("%02d/%02d", cursor+1, len(m.board.Topics)),
				fmt.Sprintf("Page: %d%s", m.page, lastPage),
				preview,
			),
		),
	)
}
