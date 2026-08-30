package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CTread",
	Short: "reading cli",
}

func Execute() {
	rootCmd.Execute()
}
