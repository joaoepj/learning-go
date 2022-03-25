package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "cobra",
  Short: "Cobra is a test using cobra library",
  Long: `A Simple library to support command switches and flags`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
    //fmt.Println("Printing args and cmd: ", args, cmd)
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}


//func init() {
//  
//}