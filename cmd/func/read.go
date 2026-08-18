package functionality

import (
	"ctread/books/importers"
	"ctread/books/reader"
	"fmt"
	"os"
	"strconv"
)

func ReadFunc(args []string) {
	if len(args) < 1 {
		fmt.Println("You must provide a book.")
		return
	}

	bookPath := args[0]

	if _, err := os.Stat(bookPath); os.IsNotExist(err) {
		bookInfo, found := importers.FindBook(bookPath)
		if !found {
			fmt.Printf("Book %q was not found.\n", bookPath)
			return
		}

		bookPath = bookInfo.Path
	}

	var content string
	var err error

	if len(args) == 1 {
		content, err = reader.Retrieve(bookPath)
	} else {
		chapter, parseErr := strconv.Atoi(args[1])
		if parseErr != nil {
			fmt.Printf("Invalid chapter: %q\n", args[1])
			return
		}

		content, err = reader.Retrieve(bookPath, chapter)
	}

	if err != nil {
		fmt.Printf("Failed to retrieve content: %v\n", err)
		return
	}

	fmt.Println(content)
}
