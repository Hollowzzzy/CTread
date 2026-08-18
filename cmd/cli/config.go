package cmd

import (
	"github.com/spf13/cobra"

	funcc "ctread/cmd/func"
)

var ConfigCmd = &cobra.Command{
	Use:   "config [setting] [value]",
	Short: "Edit the config of CTread.",
	Args:  cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		funcc.ConfigFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
}
