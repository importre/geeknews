package board

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"importre.com/geeknews/geeknews"
	"importre.com/geeknews/programs/topic"
)

type (
	FetchBoardResponse struct {
		page  uint
		board geeknews.Board
	}

	Model struct {
		board         geeknews.Board
		table         table.Model
		topic         topic.Model
		boards        list.Model
		loading       bool
		message       string
		width, height int
		page          uint
		title         string
	}

	BoardItem struct {
		title, description string
		Value              string
	}
)
