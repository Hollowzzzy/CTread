package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"ctread/config"
)

var ConfigCmd = &cobra.Command{
	Use:   "config [setting] [value]",
	Short: "Edit the config of CTread.",
	Args:  cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if args[0] != "reset" {
			if len(args) == 1 {
				err = config.Setting(args[0])
			} else {
				err = config.Setting(args[0], args[1])
			}
		} else {
			err = config.Reset()
		}
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
}
