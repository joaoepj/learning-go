// learning-go is a repository showing how to apply development
// best practices in your code while learning a programming language
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var language string

	app := &cli.App{
		Name: "learning-go",
		Usage: `a repository demonstrating how to apply development
		best practices in your code while learning the go programming language.
		For example it uses cli library to allow you to invoke different modules
		where you can test language features and implement test scenarios`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "someone"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "problems",
				Aliases: []string{"p"},
				Usage:   "A set of problem solutions implmented in golang.",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(append(os.Args, "help"))
	if err != nil {
		log.Fatal(err)
	}
}
