package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"ctread/books/importers"
	funcc "ctread/cmd/func"
)

var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage the book registry",
	Run: func(cmd *cobra.Command, args []string) {
		registry, err := importers.Load()
		if err != nil {
			fmt.Println("Failed to load registry:", err)
			return
		} else {
			fmt.Println(registry)
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
