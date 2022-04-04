package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Long:  `This is a test only command`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Testing...")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
