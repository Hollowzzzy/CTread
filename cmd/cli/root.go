package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	configg "ctread/config"
)

var rootCmd = &cobra.Command{
	Use:   "CTread",
	Short: "reading cli and tui",

	Run: func(cmd *cobra.Command, args []string) {
		config, err := configg.Load()
		if err != nil {
			config = configg.Config(configg.Defaults)
		}
		tui := config.TuiEnabled
		if tui {
			fmt.Println("Tui is enabled!")
		} else {
			fmt.Println(cmd.Help())
		}
	},
}

func Execute() {
	rootCmd.Execute()
}
