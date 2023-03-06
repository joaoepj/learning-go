package lg_problems

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var TwoSum_arr = []int{2, 7, 11, 15}

// This problem is actually from Leet Code
var SCTwoSum *cli.Command = &cli.Command{
	Name:    "twosum",
	Aliases: []string{"tw"},
	Usage:   "Two Sum Problem",
	Action: func(c *cli.Context) error {
		TwoSum(TwoSum_arr, 9)

		return nil
	},
}

func TwoSum(nums []int, target int) []int {
	var L, R int = 0, 1
	var result []int

	for {

		fmt.Println(nums)
		if shoot := nums[L] + nums[R]; shoot == target {
			fmt.Println(shoot)
			result = append(result, L)
			result = append(result, R)
			break
		} else if R < len(nums)-1 {
			R++

		} else {
			L++
			R = L + 1
		}

		fmt.Println(L, R)
	}
	fmt.Println(result)
	return result

}
