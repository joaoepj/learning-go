package lg_misc

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var SCRandomString *cli.Command = &cli.Command{
	Name:    "randomstring",
	Aliases: []string{"rs"},
	Usage:   "Random String Generator",
	Action: func(c *cli.Context) error {
		size, err := strconv.Atoi(c.Args().First())
		if err != nil {
			os.Exit(1)
		}
		for i := 0; i < size; i++ {
			fmt.Println(RandomString(size))
		}

		return nil
	},
}

// Executes when module is imported
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(size int) string {
	buffer := make([]rune, size)
	for i := range buffer {
		buffer[i] = letters[rand.Intn(len(letters))]
	}
	return string(buffer)
}

var SCRandomIntegers *cli.Command = &cli.Command{
	Name:    "randominteger",
	Aliases: []string{"ri"},
	Usage:   "Random Integer Generator",
	Action: func(c *cli.Context) error {
		size, err := strconv.Atoi(c.Args().First())
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(RandomIntegers(size))

		return nil
	},
}

func RandomIntegers(size int) []int {
	integers := make([]int, 0)
	for i := 0; i < size; i++ {
		integers = append(integers, rand.Intn(255))
	}
	return integers
}
