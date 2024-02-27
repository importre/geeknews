package geeknews

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"importre.com/geeknews/components"
)

type Topic struct {
	Id       uint
	VoteEnum uint
	Url      string
	Title    string
	Preview  string
	Content  string
	Info     TopicInfo
	Comments []Comment
}

type TopicInfo struct {
	Points      uint
	User        string
	Timestamp   string
	NumComments uint
}

func (topic Topic) Markdown(width int) string {
	return components.Markdown(topic.Content, width, topic.Title)
}

type Comment struct {
	Id    string
	Depth uint
	Info  CommentInfo
}

type CommentInfo struct {
	User      string
	Timestamp string
	Content   string
}

func (comment *Comment) Markdown(width int) string {
	return components.Markdown(comment.Info.Content, width, "")
}

type TopicRequest struct {
	Id uint
}

type TopicResponse struct {
	Topic Topic
}

func FetchTopic(request TopicRequest) (*TopicResponse, error) {
	url := fmt.Sprintf("%s/topic?id=%d", BaseUrl, request.Id)
	doc, error := get(url)
	if error != nil {
		return nil, error
	}

	topicElement := doc.Find("body > main > article > div.topic-table > div.topic")
	titleElement := topicElement.Find("div.topictitle > a")
	topicInfoElement := topicElement.Find("div.topicinfo")
	title := titleElement.Find("h1").Text()
	content, _ := topicElement.Find("div.topic_contents").Html()

	comments := []Comment{}
	doc.
		Find("body > main > article > div.comment_thread > div.comment_row").
		Each(func(_ int, s *goquery.Selection) {
			commentInfoElement := s.Find("div.commentinfo")
			commentContent, _ := s.Find("div.commentTD").Html()
			comment := Comment{
				Id:    s.AttrOr("id", ""),
				Depth: parseUint(regexDepth.FindStringSubmatch(s.AttrOr("style", ""))[1]),
				Info: CommentInfo{
					User:      commentInfoElement.Find("a:nth-child(1)").Text(),
					Timestamp: commentInfoElement.Find("a:nth-child(2)").Text(),
					Content:   commentContent,
				},
			}
			comments = append(comments, comment)
		})

	topic := &TopicResponse{
		Topic{
			Id:       request.Id,
			VoteEnum: 0,
			Url:      url,
			Title:    title,
			Content:  content,
			Info:     parseTopicInfo(topicInfoElement),
			Comments: comments,
		},
	}
	return topic, nil
}

func parseTopicInfo(topicInfoElement *goquery.Selection) TopicInfo {
	userElement := topicInfoElement.Find("a").First()
	commentsMatch := regexComments.FindStringSubmatch(topicInfoElement.Find("a").Last().Text())
	var numComments uint
	if len(commentsMatch) == 2 {
		numComments = parseUint(commentsMatch[1])
	}
	return TopicInfo{
		Points:      parseUint(topicInfoElement.Find("span:nth-child(1)").Text()),
		User:        strings.TrimSpace(userElement.Text()),
		Timestamp:   strings.TrimSpace(userElement.Get(0).NextSibling.Data),
		NumComments: numComments,
	}
}

var (
	regexComments = regexp.MustCompile(`^댓글 (\d+)개$`)
	regexDepth    = regexp.MustCompile(`^--depth:(\d+)$`)
)
