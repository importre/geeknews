package board

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"importre.com/geeknews/programs/topic"
)

func New() Model {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	boards := list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		0,
		0,
	)
	boards.Title = "글 목록"
	boards.SetFilteringEnabled(false)

	return Model{
		table: table.New(
			table.WithFocused(true),
			table.WithStyles(s),
		),
		topic:   topic.New(),
		boards:  boards,
		loading: true,
		page:    1,
		title:   boardItems[0].title,
	}
}

func (boardItem BoardItem) Title() string {
	return boardItem.title
}

func (boardItem BoardItem) Description() string {
	return boardItem.description
}

func (boardItem BoardItem) FilterValue() string {
	return boardItem.title
}
