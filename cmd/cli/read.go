package cmd

import (
	"github.com/spf13/cobra"

	functionality "ctread/cmd/func"
)

var ReadCmd = &cobra.Command{
	Use:   "read [book] [chapter]",
	Short: "Read a book or a specific chapter.",

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		functionality.ReadFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(ReadCmd)
}
