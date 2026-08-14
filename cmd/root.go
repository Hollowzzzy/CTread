package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "CTread",
	Short: "reading cli and tui",
}

func Execute() {
	rootCmd.Execute()
}
