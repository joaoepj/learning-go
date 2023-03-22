package lg_misc

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var SCRemoveVowels *cli.Command = &cli.Command{
	Name:    "removevowels",
	Aliases: []string{"rw"},
	Usage:   "Remove Vowels",
	Action: func(c *cli.Context) error {

		s := c.Args().First()

		if s == "" {
			s = RandomString(100)
		}
		fmt.Println(c.Command.FullName())
		fmt.Println(RemoveVowels(s))

		return nil
	},
}

func RemoveVowels(s string) string {
	b := make([]byte, len(s))
	for _, c := range s {
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' && c != 'A' && c != 'E' && c != 'I' && c != 'O' && c != 'U' {
			b = append(b, byte(c))
		}
	}
	//fmt.Println(string(b))
	return string(b)
}
