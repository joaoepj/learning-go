package lg_problems

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var CountSwaps_arr = []int32{5, 3, 4, 2, 1}

var SCCountSwaps *cli.Command = &cli.Command{
	Name:    "countswaps",
	Aliases: []string{"cs"},
	Usage:   "Count Swaps Problem",
	Action: func(c *cli.Context) error {
		CountSwaps(CountSwaps_arr)

		return nil
	},
}

func swap1(x int32, y int32) (int32, int32) {
	return y, x
}

func CountSwaps(a []int32) {
	// Write your code here
	var numSwaps int
	//fmt.Println(numSwaps, a)
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = swap1(a[j], a[j+1])
				numSwaps += 1
				//fmt.Println(numSwaps, a)
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numSwaps)
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[len(a)-1])
}
