package topic

import (
	"github.com/charmbracelet/bubbles/viewport"
	"importre.com/geeknews/geeknews"
)

type (
	FetchTopicResponse struct {
		topic geeknews.Topic
	}

	Loading struct{}

	Model struct {
		width, height int
		loading       bool
		topic         geeknews.Topic
		viewport      viewport.Model
	}
)
