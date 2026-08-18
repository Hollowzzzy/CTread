package cmd

import (
	funcc "ctread/cmd/func"

	"github.com/spf13/cobra"
)

var Importcmd = &cobra.Command{
	Use:   "import [file path...]",
	Short: "Import one or more books and add them to the registry. Epub is the only supported format as of now.",

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		funcc.ImportFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(Importcmd)
}
