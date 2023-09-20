package lg_problems

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var UtopianTree_arr = [][]int32{
	{4, 6, 5, 3, 3, 1},
	{1, 1, 2, 2, 4, 4, 5, 5, 5},
}

var SCUtopianTree *cli.Command = &cli.Command{
	Name:    "utopiantree",
	Aliases: []string{"ut"},
	Usage:   "Utopian Tree Problem",
	Action: func(c *cli.Context) error {
		fmt.Println(utopianTree(7))
		return nil
	},
}

func utopianTree(n int32) int32 {
	// Write your code here
	var height int32 = 1
	for i := int32(0); i < n; i++ {
		//fmt.Println("pre","i",i,"height",height)
		if i%2 == 0 {
			height = height * 2
		} else {
			height++
		}
		//fmt.Println("pos","i",i,"height",height)
	}
	return height

}
