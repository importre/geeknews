package geeknews_test

import (
	"testing"

	"importre.com/geeknews/geeknews"
)

func TestFetchBoard(t *testing.T) {
	request := geeknews.BoardRequest{Key: "", Page: 1}
	response, error := geeknews.FetchBoard(request)

	if error != nil {
		t.Error(error)
	}

	if len(response.Board.Topics) < 1 {
		t.Error("게시글이 없음")
	}

	for _, topic := range response.Board.Topics {
		t.Log(topic)
	}
}

func TestFetchTopic(t *testing.T) {
	request := geeknews.TopicRequest{Id: 13581}
	response, error := geeknews.FetchTopic(request)

	if error != nil {
		t.Error(error)
	}

	t.Log(response.Topic)
}
