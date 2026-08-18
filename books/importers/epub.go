package importers

import (
	"fmt"
	"os"

	"github.com/simp-lee/epub"
)

func Epub(path string, returnn bool) (*epub.Book, bool) {
	book, err := epub.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	if !returnn {
		defer book.Close()
	}

	md := book.Metadata()

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		book.Close()
		return nil, false
	}

	modified := info.ModTime()

	RegistryAdd(md.Titles[0], path, modified)

	return book, returnn
}
