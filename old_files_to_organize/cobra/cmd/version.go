package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringVar(&Version, "token", "", "Specify the token")
	//rootCmd.PersistentFlags().StringVarP(&Version, "token", "t", "", "version output")
  }
  
  var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cobra",
	Long:  `All software has versions. This is Cobra's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args, cmd.Flag("token"))
	  fmt.Println("Cobra library tester v0.1")
	},
  }