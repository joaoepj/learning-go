package lg_problems

import (
	"fmt"
	"strconv"

	"github.com/joaoepj/learning-go/lg_misc"
	"github.com/urfave/cli/v2"
)

var BeautifulDays_arr = [][]int32{
	{4, 6, 5, 3, 3, 1},
	{1, 1, 2, 2, 4, 4, 5, 5, 5},
}

var SCBeautifulDays *cli.Command = &cli.Command{
	Name:    "beaultifuldays",
	Aliases: []string{"bd"},
	Usage:   "Beaultiful Days Problem",
	Action: func(c *cli.Context) error {
		fmt.Println(beautifulDays(3, 5, 7))
		return nil
	},
}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var bDays int32
	for x := i; x <= j; x++ {
		if lg_misc.Abs32(x-reverse(x))%k == 0 {
			bDays++
		}
	}

	return bDays
}

func reverse(n int32) int32 {
	str := strconv.Itoa(int(n))
	var rev string
	for _, v := range str {
		rev = string(v) + rev
	}

	i, _ := strconv.Atoi(rev)
	return int32(i)
}
