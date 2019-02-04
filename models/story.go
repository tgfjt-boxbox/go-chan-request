package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	stories = "https://hacker-news.firebaseio.com/v0/topstories.json"
	item    = "https://hacker-news.firebaseio.com/v0/item/%d.json"
)

type Story struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func (story *Story) EchoTitle() {
	fmt.Println(fmt.Sprintf("「%s」", story.Title))
}

func GetTopStories(c *http.Client) ([]uint64, error) {
	var err error
	var storyIds []uint64

	req, _ := http.NewRequest("GET", stories, nil)

	res, err := c.Do(req)

	fmt.Println("Requested.")

	if err != nil {
		return storyIds, errors.WithStack(err)
	}

	err = json.NewDecoder(res.Body).Decode(&storyIds)

	if err != nil {
		return storyIds, errors.WithStack(err)
	}

	return storyIds, nil
}

func GetStory(c *http.Client, sID uint64) (Story, error) {
	var err error
	var story Story

	req, _ := http.NewRequest("GET", fmt.Sprintf(item, sID), nil)
	res, err := c.Do(req)

	if err != nil {
		return story, errors.WithStack(err)
	}

	err = json.NewDecoder(res.Body).Decode(&story)

	if err != nil {
		return story, errors.WithStack(err)
	}

	return story, nil
}
