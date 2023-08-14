package lg_algorithms

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var SCFibonacci *cli.Command = &cli.Command{
	Name:    "dynamicprogramming",
	Aliases: []string{"dp"},
	Usage:   "Dynamic Programming",
	Action: func(c *cli.Context) error {
		// call CacheFibonacci first so that is ease to see
		// how it answers much faster than RecFibonacci
		fmt.Println("CacheFibonacci:", CacheFibonacci(cfMap, 40))
		fmt.Println("RecFibonacci:", RecFibonacci(40))
		return nil
	},
}

func RecFibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return RecFibonacci(n-1) + RecFibonacci(n-2)

}

var cfMap map[int]int

func cacheFibonacci(cfMap map[int]int, n int) int {
	if _, ok := cfMap[n]; !ok {
		cfMap[n] = cacheFibonacci(cfMap, n-1) + cacheFibonacci(cfMap, n-2)
	}

	return cfMap[n]
}

func CacheFibonacci(cfMap map[int]int, n int) int {
	cfMap = make(map[int]int)
	cfMap[0] = 0
	cfMap[1] = 1

	return cacheFibonacci(cfMap, n)
}
