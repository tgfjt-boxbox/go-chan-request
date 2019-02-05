package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tgfjt-boxbox/go-chan-request/models"
)

func SaveLinks(ids []uint, list map[uint]*models.Story) {
	var writer *bufio.Writer
	t := time.Now()
	const layout = "2006-01-02-15-04-05"

	p, err := filepath.Abs(fmt.Sprintf("result/%s.md", t.Format(layout)))
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0644)
	writer = bufio.NewWriter(file)

	// in order of topstories
	for _, id := range ids {
		story := list[id]
		writer.WriteString(fmt.Sprintf("- [%s](%s)", story.Title, story.Url) + "\n")
	}

	writer.Flush()
}
