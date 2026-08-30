package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"ctread/books/importers"
	funcc "ctread/cmd/func"
	Styles "ctread/lipgloss"
)

var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage the book registry",
	Run: func(cmd *cobra.Command, args []string) {
		registry, err := importers.Load()
		if err != nil {
			fmt.Printf("%s\n", Styles.ERR.Render(fmt.Sprintf("Failed to load registry: %v", err)))
			return
		}

		if len(registry) == 0 {
			fmt.Println(Styles.INFO.Render("The registry is empty."))
			return
		}

		for i, book := range registry {
			fmt.Printf("%d. %s\n", i+1, Styles.INFO.Render(book.Name))
			fmt.Printf("   Path: %s\n", Styles.SUCCESS.Render(book.Path))
			fmt.Printf("   Modified: %s\n\n", Styles.SUCCESS.Render(book.Modified.Format("2006-01-02 15:04:05")))
		}
	},
}

var RegistryAddCmd = &cobra.Command{
	Use:   "add [path]",
	Short: "Add a book to the registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		funcc.ImportFunc(args)
	},
}

var RegistryDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Remove a book from the registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		importers.RegistryDelete(args[0])
	},
}

var RegistryClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the entire registry",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		importers.RegistryClear()
	},
}

func init() {
	RegistryCmd.AddCommand(
		RegistryAddCmd,
		RegistryDeleteCmd,
		RegistryClearCmd,
	)

	rootCmd.AddCommand(RegistryCmd)
}
