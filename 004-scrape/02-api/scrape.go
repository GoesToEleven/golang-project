package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type tweet struct {
	Name     string
	Username string
	Message  string
}

type conversationResponse struct {
	MinPos string `json:"min_position"`
	More   bool   `json:"has_more_items"`
	Html   string `json:"items_html"`
}

func makeConversationRequest(maxPos string) (*conversationResponse, error) {
	params := url.Values{}
	params.Set("include_available_features", "1")
	params.Set("include_entities", "1")
	params.Set("max_position", maxPos)
	params.Set("reset_error_state", "false")

	url := "https://twitter.com/i/Todd_McLeod/conversation/1169751640926146560?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error while getting conversation data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Incorrect http status code: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	cr := &conversationResponse{}
	err = json.NewDecoder(resp.Body).Decode(cr)
	if err != nil {
		return nil, fmt.Errorf("Error while decoding response: %w", err)
	}

	return cr, nil
}

func getConversation() ([]string, error) {
	continueCode := ""
	messages := []string{}

	for {
		resp, err := makeConversationRequest(continueCode)
		if err != nil {
			return nil, fmt.Errorf("Unable to make conversation request: %w", err)
		}

		messages = append(messages, resp.Html)

		if !resp.More {
			break
		}

		continueCode = resp.MinPos
		time.Sleep(time.Second)
	}

	return messages, nil
}

func parseHtml(message string) ([]tweet, error) {
	rdr := strings.NewReader(message)
	doc, err := goquery.NewDocumentFromReader(rdr)
	if err != nil {
		return nil, fmt.Errorf("Unable to read html: %w", err)
	}

	ts := []tweet{}

	doc.Find(".tweet").Each(func(i int, s *goquery.Selection) {
		ts = append(ts, tweet{
			Name:     s.Find(".account-group .fullname").Text(),
			Username: s.Find(".account-group .username").Text(),
			Message:  s.Find(".tweet-text").Text(),
		})
	})

	return ts, nil
}

func main() {
	resp, err := getConversation()
	if err != nil {
		panic(err)
	}

	tweets := []tweet{}
	for _, msg := range resp {
		ts, err := parseHtml(msg)
		if err != nil {
			panic(err)
		}
		tweets = append(tweets, ts...)
	}

	bs, err := json.MarshalIndent(tweets, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of tweets:", len(tweets))
}
