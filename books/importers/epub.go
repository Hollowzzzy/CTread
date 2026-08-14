package importers

import (
	"fmt"
	"log"

	"github.com/simp-lee/epub"
)

func import_book() {
	book, err := epub.Open("book.epub")
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

	// Metadata
	md := book.Metadata()
	if len(md.Titles) > 0 {
		fmt.Println("Title:", md.Titles[0])
	}
	for _, a := range md.Authors {
		fmt.Println("Author:", a.Name)
	}

	// Table of Contents
	for _, item := range book.TOC() {
		fmt.Printf("  %s → %s\n", item.Title, item.Href)
	}

	// Chapters
	for _, ch := range book.Chapters() {
		text, err := ch.TextContent()
		if err != nil {
			continue
		}
		fmt.Printf("  [%s] %d chars\n", ch.Title, len(text))
	}

	// Cover image
	cover, err := book.Cover()
	if err == nil {
		fmt.Printf("Cover: %s (%d bytes)\n", cover.MediaType, len(cover.Data))
	}
}
