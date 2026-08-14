package cmd

import (
	"github.com/spf13/cobra"

	"ctread/config"
)

var ConfigCmd = &cobra.Command{
	Use:   "config [setting] [value]",
	Short: "Edit the config of CTread.",

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		setting := args[0]

		if len(args) == 1 {
			config.UpdateSetting(setting)
			return
		}

		config.UpdateSetting(setting, args[1])
	},
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
}
