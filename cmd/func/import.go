package functionality

import (
	"ctread/books/importers"
	"fmt"
	"path/filepath"
)

func ImportFunc(args []string, code ...bool) {
	showCode := false

	if len(code) > 0 {
		showCode = code[0]
	}

	for _, path := range args {
		fileType := GetFileType(path)

		switch fileType {
		case "epub":
			if showCode {
				fmt.Printf("Importing %s...\n", path)
			} else {
				fmt.Printf("Adding %s to registry...\n", path)
			}

			importers.Epub(path, false)

		default:
			fmt.Printf("I don't know what %s is.\n", path)
		}
	}
}

func GetFileType(path string) string {
	extension := filepath.Ext(path)

	if len(extension) > 0 {
		return extension[1:]
	}

	return ""
}
