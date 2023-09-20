package lg_problems

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var ViralAdvertising_arr = [][]int32{
	{4, 6, 5, 3, 3, 1},
	{1, 1, 2, 2, 4, 4, 5, 5, 5},
}

var SCViralAdvertisement *cli.Command = &cli.Command{
	Name:    "viraladvertisement",
	Aliases: []string{"va"},
	Usage:   "Viral Advertisement Problem",
	Action: func(c *cli.Context) error {
		fmt.Println(viralAdvertising(7))
		return nil
	},
}

func viralAdvertising(n int32) int32 {
	// Write your code here
	var shared, liked, cumulative int32 = 5, 0, 0
	for day := int32(1); day <= n; day++ {
		liked = shared / 2
		cumulative += liked
		shared = liked * 3
	}

	return cumulative
}
