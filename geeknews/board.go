package geeknews

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Board struct {
	Topics     []Topic
	Key        string
	IsLastPage bool
}

type BoardRequest struct {
	Key  string
	Page uint
}

type BoardResponse struct {
	Board Board
}

func FetchBoard(request BoardRequest) (*BoardResponse, error) {
	url := fmt.Sprintf("%s/%s?page=%d", BaseUrl, request.Key, request.Page)
	doc, error := get(url)
	if error != nil {
		return nil, error
	}

	topics := []Topic{}
	regexTopicId := regexp.MustCompile(`id=(\d+)`)
	doc.
		Find("body > main > article > div.topics > div.topic_row").
		Each(func(_ int, s *goquery.Selection) {
			topicDescElement := s.Find("div.topicdesc > a")
			topicInfoElement := s.Find("div.topicinfo").First()
			topicInfo := parseTopicInfo(topicInfoElement)
			topicId := regexTopicId.FindStringSubmatch(topicDescElement.AttrOr("href", "id=0"))[1]
			topic := Topic{
				Id:       parseUint(topicId),
				VoteEnum: parseUint(s.Find("div.votenum").Text()),
				Url:      fmt.Sprintf("%s/topic?id=%s", BaseUrl, topicId),
				Title:    strings.TrimSpace(s.Find("div.topictitle > a").Text()),
				Preview:  strings.TrimSpace(s.Find("div.topicdesc > a").Text()),
				Info:     topicInfo,
			}
			topics = append(topics, topic)
		})

	isLastPage := doc.Find("body > main > article > div > div.next.commentTD").Length() == 0

	board := Board{
		Topics:     topics,
		Key:        request.Key,
		IsLastPage: isLastPage,
	}

	boards := &BoardResponse{
		Board: board,
	}

	return boards, nil
}
