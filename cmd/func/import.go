package functionality

import (
	"fmt"
	"path/filepath"

	"github.com/Hollowzzzy/CTread/books/importers"

	Styles "github.com/Hollowzzzy/CTread/lipgloss"
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
				fmt.Println(Styles.INFO.Render(fmt.Sprintf("Importing %s...", path)))
			} else {
				fmt.Println(Styles.INFO.Render(fmt.Sprintf("Adding %s to registry...", path)))
			}

			importers.Epub(path, false)

		default:
			fmt.Println(Styles.ERR.Render(fmt.Sprintf(
				"I don't know what %s is...",
				path,
			)))
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
