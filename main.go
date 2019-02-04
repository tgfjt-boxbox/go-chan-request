package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/tgfjt-boxbox/go-chan-request/client"
	"github.com/tgfjt-boxbox/go-chan-request/models"
)

func init() {
	fmt.Println("Init.")
}

func main() {
	var err error
	c := client.GetClient()
	storyIds, err := models.GetTopStories(c)

	if err != nil {
		log.Fatalf("fail to get story Ids!")
	}

	fmt.Println("Got StoryIds")

	var wg sync.WaitGroup

	getStory := func(sID uint64) {
		story, err := models.GetStory(c, sID)

		if err != nil {
			log.Fatalf("fail to get story")
		}

		story.EchoTitle()
		wg.Done()
	}

	for _, sID := range storyIds {
		wg.Add(1)
		// goroutine
		go getStory(sID)
	}

	wg.Wait()

	fmt.Println("All Done!!")
}
