package geeknews

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func get(url string) (*goquery.Document, error) {
	resp, error := http.Get(url)
	if error != nil {
		return nil, error
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		message := fmt.Sprintf("status code: %d", resp.StatusCode)
		return nil, &HttpError{resp.StatusCode, message}
	}

	doc, error := goquery.NewDocumentFromReader(resp.Body)
	if error != nil {
		return nil, error
	}

	return doc, nil
}

type HttpError struct {
	statusCode int
	message    string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("[%d]: %s", e.statusCode, e.message)
}
