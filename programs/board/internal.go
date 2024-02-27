package board

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) setTableLayout() {
	pointsLength := 0
	for _, topic := range m.board.Topics {
		pointsLength = max(pointsLength, len(fmt.Sprintf("%4d", topic.Info.Points)))
	}
	commentsLength := 0
	for _, topic := range m.board.Topics {
		commentsLength = max(commentsLength, len(fmt.Sprintf("%4d", topic.Info.NumComments)))
	}
	userLength := 0
	for _, topic := range m.board.Topics {
		userLength = max(userLength, len(topic.Info.User)+2)
	}
	timeLength := 0
	for _, topic := range m.board.Topics {
		timeLength = max(timeLength, len(topic.Info.Timestamp)+2)
	}
	columns := []table.Column{
		{Title: "포인트", Width: max(pointsLength, 6)},
		{Title: "타이틀", Width: m.width},
		{Title: "댓글", Width: max(commentsLength, 4)},
		{Title: "작성자", Width: userLength},
		{Title: "시간", Width: timeLength},
	}

	titleIndex := 0
	for i, column := range columns {
		if columns[titleIndex].Width < column.Width {
			titleIndex = i
		}
	}

	columns[titleIndex].Width -= screenStyle.GetHorizontalFrameSize()

	for i, c := range columns {
		if i != titleIndex {
			columns[titleIndex].Width -= c.Width
		}
		columns[titleIndex].Width -= 2
	}

	m.table.SetColumns(columns)
	m.setTableHeight()
}

func (m *Model) setTableHeight() {
	tableHeight := m.height - screenStyle.GetVerticalFrameSize() - 1
	m.table.SetHeight(tableHeight)
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

func (m Model) visibleBoards() bool {
	return len(m.boards.Items()) > 0
}

func (m Model) update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.boards.SetItems([]list.Item{})
			return m, cmd

		case "enter":
			boardItem := boardItems[m.boards.Index()]
			m.table.SetCursor(0)
			m.title = boardItem.title
			return m, m.Init(boardItem.Value, 1)

		case "esc":
			return m, cmd

		default:
			m.boards, cmd = m.boards.Update(msg)
			return m, cmd
		}
	}

	return m, cmd
}

var (
	screenStyle = lipgloss.NewStyle().
			Margin(2, 8)

	boardItems = []BoardItem{
		{title: "GeekNews", description: "긱뉴스 홈", Value: ""},
		{title: "최신글", description: "최신글 목록", Value: "new"},
		{title: "Ask", description: "질문 글 목록", Value: "ask"},
		{title: "Show", description: "직접 만든 오픈소스나 개발한 서비스 홍보 글 목록", Value: "show"},
		{title: "GN+", description: "AI가 요약한 최신 뉴스 목록", Value: "plus"},
		{title: "인기글", description: "최근 6개월간 upvote 를 많이 받은 글 목록", Value: "lists/popular"},
		{title: "인기글 (전체)", description: "전체에서 upvote 를 많이 받은 글 목록", Value: "lists/all-time-popular"},
		{title: "읽을만한 긴 글", description: "최근 6개월간 읽을만한 긴 글 목록", Value: "lists/long-read"},
		{title: "읽을만한 긴 글 (전체)", description: "전체에서 읽을만한 긴 글 목록", Value: "lists/all-time-long-read"},
		{title: "활발한 글", description: "최근 6개월간 댓글이 활발한 글 목록", Value: "lists/active"},
		{title: "활발한 글 (전체)", description: "전체에서 댓글이 활발한 글 목록", Value: "lists/all-time-active"},
	}
)
