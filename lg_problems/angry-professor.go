package lg_problems

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var AngryProfessor_arr = [][]int32{
	{4, 6, 5, 3, 3, 1},
	{1, 1, 2, 2, 4, 4, 5, 5, 5},
}

var SCAngryProfessor *cli.Command = &cli.Command{
	Name:    "angryprofessor",
	Aliases: []string{"ap"},
	Usage:   "Angry Professor Problem",
	Action: func(c *cli.Context) error {
		fmt.Println(angryProfessor(3, AngryProfessor_arr[0]))
		return nil
	},
}

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var earliers int32
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			earliers++
		}
	}

	if earliers >= k {
		return "NO"
	} else {
		return "YES"
	}

}
