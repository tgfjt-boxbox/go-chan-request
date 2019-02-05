package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/tgfjt-boxbox/go-chan-request/models"
	client "github.com/tgfjt-boxbox/go-chan-request/utils"
	markdown "github.com/tgfjt-boxbox/go-chan-request/utils"
)

var mu sync.RWMutex
var wg sync.WaitGroup
var err error

func main() {
	c := client.GetClient()
	storyIds, err := models.GetTopStories(c)

	mu = sync.RWMutex{}

	if err != nil {
		log.Fatalf("fail to get story Ids!")
	}

	fmt.Println("got topstories IDs successfully.")

	var storiesMap = make(map[uint]*models.Story, len(storyIds))

	getStory := func(sID uint, list map[uint]*models.Story) {
		defer wg.Done()
		story, err := models.GetStory(c, sID)

		if err != nil {
			log.Fatalf("fail to get story.")
		}

		mu.Lock()
		defer mu.Unlock()

		list[sID] = &story
	}

	for _, sID := range storyIds {
		wg.Add(1)
		// goroutine
		go getStory(sID, storiesMap)
	}

	wg.Wait()

	markdown.SaveLinks(storyIds, storiesMap)
}
