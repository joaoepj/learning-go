// learning-go is a repository showing how to apply development
// best practices in your code while learning a programming language
package main

import (
	"log"
	"os"

	problems "github.com/joaoepj/learning-go/lg_problems"
	snmp "github.com/joaoepj/learning-go/lg_snmp"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name: "learning-go",

		Usage: `a repository demonstrating how to apply development
		best practices in your code while learning the go programming language.
		For example it uses cli library to allow you to invoke different modules
		where you can test language features and implement test scenarios`,
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "problems",
				Aliases: []string{"p"},
				Usage:   "A set of problem solutions implmented in golang.",
				Subcommands: []*cli.Command{
					{
						Name:    "staircase",
						Aliases: []string{"s"},
						Usage:   "Staircase problem",
						Action: func(c *cli.Context) error {
							problems.Staircase(5)
							return nil
						},
					},
					{
						Name:    "minmaxsum",
						Aliases: []string{"m"},
						Usage:   "Min-Max Sum problem",
						Action: func(c *cli.Context) error {
							arr := []int32{2, 8, 5, 9, 2, 5, 7, 1}
							problems.MinMaxSum(arr)
							return nil
						},
					},
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
			{
				Name:    "snmp",
				Aliases: []string{"s"},
				Usage:   "Test the SNMP library gosnmp",
				Subcommands: []*cli.Command{
					{
						Name:    "snmpget",
						Aliases: []string{"g"},
						Usage:   "snmpget",
						Action: func(c *cli.Context) error {
							snmp.SnmpGet()
							return nil
						},
					},
					{
						Name:    "snmpwalk",
						Aliases: []string{"w"},
						Usage:   "snmpwalt",
						Action: func(c *cli.Context) error {
							snmp.SnmpWalk()
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(append(os.Args, "help"))
	if err != nil {
		log.Fatal(err)
	}
}
