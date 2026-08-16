package importers

import (
	"fmt"
	"log"
	"os"

	"github.com/simp-lee/epub"
)

func Epub(path string) {
	book, err := epub.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

	md := book.Metadata()

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	modified := info.ModTime()

	RegistryAdd(md.Titles[0], path, modified)
}
