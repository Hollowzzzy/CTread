package reader

import (
	"fmt"

	epub "github.com/simp-lee/epub"
)

func Retrieve(path string, chapters ...int) (string, error) {
	book, err := epub.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open book: %w", err)
	}
	defer book.Close()

	contentChapters := book.ContentChapters()

	if len(contentChapters) <= 1 {
		return "", fmt.Errorf("book contains no readable chapters")
	}

	chapter := 1

	if len(chapters) > 0 {
		chapter = chapters[0]
	}

	if chapter < 1 || chapter >= len(contentChapters) {
		return "", fmt.Errorf("chapter %d does not exist", chapter)
	}

	content, err := contentChapters[chapter].TextContent()
	if err != nil {
		return "", fmt.Errorf("failed to read chapter %d: %w", chapter, err)
	}

	return content, nil
}
