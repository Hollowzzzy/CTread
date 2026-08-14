package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ReadCmd = &cobra.Command{
	Use:   "read [book] [chapter] [page]",
	Short: "Read a book whether it be a chapter, a page, or more.",

	Args: cobra.ArbitraryArgs,

	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			book := args[0]
			fmt.Printf("reading %s.\n", book)

		case 2:
			book := args[0]
			chapter := args[1]
			fmt.Printf("reading chapter %s of %s.\n", chapter, book)

		case 3:
			book := args[0]
			chapter := args[1]
			page := args[2]
			fmt.Printf("reading page %s of chapter %s of %s.\n", page, chapter, book)

		default:
			if len(args) < 1 {
				fmt.Println("You must provide at least a book.")
			} else if len(args) > 3 {
				fmt.Println("Too many arguments. Usage: read [book] [chapter] [page]")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ReadCmd)
}
