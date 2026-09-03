package functionality

import (
	"fmt"
	"os"
	"strconv"

	Styles "github.com/Hollowzzzy/CTread/lipgloss"

	"github.com/Hollowzzzy/CTread/books/reader"

	"github.com/Hollowzzzy/CTread/books/importers"
)

func ReadFunc(args []string) {
	if len(args) < 1 {
		fmt.Println(Styles.ERR.Render("You must provide a book."))
		return
	}

	bookPath := args[0]

	if _, err := os.Stat(bookPath); os.IsNotExist(err) {
		bookInfo, found := importers.FindBook(bookPath)
		if !found {
			fmt.Printf("%s\n", Styles.ERR.Render(fmt.Sprintf("Book %q was not found.", bookPath)))
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
			fmt.Printf("%s\n", Styles.ERR.Render(fmt.Sprintf("Invalid chapter: %q", args[1])))
			return
		}

		content, err = reader.Retrieve(bookPath, chapter)
	}

	if err != nil {
		fmt.Printf("%s\n", Styles.ERR.Render(fmt.Sprintf("Failed to retrieve content: %v", err)))
		return
	}

	fmt.Println(content)
}
