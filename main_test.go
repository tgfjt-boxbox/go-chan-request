package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/tgfjt-boxbox/go-chan-request/client"
	"github.com/tgfjt-boxbox/go-chan-request/models"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	os.Exit(ret)
}

func TestGetTopStories(t *testing.T) {
	var err error
	var c *http.Client
	var storyIds []uint64

	c = client.GetClient()
	storyIds, err = models.GetTopStories(c)

	if err != nil {
		t.Error(err)
	}

	if len(storyIds) < 1 {
		t.Error("not enough storyIds!")
	}
}
func TestGetStory(t *testing.T) {
	var err error
	var c *http.Client
	var story models.Story

	c = client.GetClient()
	story, err = models.GetStory(c, 19064875)

	if err != nil {
		t.Error(err)
	}

	if story.Url != "https://nspk.com/" {
		t.Error("wrong url")
	}
}