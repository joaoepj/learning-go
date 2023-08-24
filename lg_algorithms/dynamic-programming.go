package lg_algorithms

import (
	"fmt"

	"github.com/joaoepj/learning-go/lg_misc"
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
		fmt.Println("DynProgFibonacci:", DynProgFibonacci(40))
		fmt.Println("UltimateFibonacci:", UltimateFibonacci(40))
		fmt.Println("RecLongestIncreasingSequence:", RecLongestIncreasingSequence("abcde", "ace", 0, 0))

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

func DynProgFibonacci(n int) int {
	dpf := make(map[int]int)

	dpf[0] = 0
	dpf[1] = 1
	for i := 2; i <= n; i++ {
		dpf[i] = dpf[i-1] + dpf[i-2]
	}

	return dpf[n]

}

func UltimateFibonacci(n int) int {
	var tmp0 int = 0
	var tmp1 int = 1
	var next int

	if n == 0 {
		return 0
	}

	for i := 2; i < n; i++ {
		next = tmp0 + tmp1
		tmp1 = tmp0
		tmp0 = next
	}

	return tmp0 + tmp1
}

func RecLongestIncreasingSequence(s1 string, s2 string, i int, j int) int {
	if i > len(s1)-1 || j > len(s2)-1 {
		return 0
	} else if s1[i] == s2[j] {
		return 1 + RecLongestIncreasingSequence(s1, s2, i+1, j+1)
	} else {
		return lg_misc.Max(RecLongestIncreasingSequence(s1, s2, i+1, j),
			RecLongestIncreasingSequence(s1, s2, i, j+1))
	}
}
