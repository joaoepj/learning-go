package lg_misc

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var SCRecursive *cli.Command = &cli.Command{
	Name:    "recursive",
	Aliases: []string{"rc"},
	Usage:   "Recursive Functions",
	Action: func(c *cli.Context) error {
		fmt.Println("Reverse: ", Reverse(PopulateInt32Slice(10), []int32{}))
		fmt.Println("Direct: ", Direct(PopulateInt32Slice(10), []int32{}))
		fmt.Println("Operator: ", Operator(10))
		fmt.Println("Operator2: ", Operator2(PopulateInt32Slice(10)))
		fmt.Println("PopulateInt32Slice: ", PopulateInt32Slice(10))
		fmt.Println("PopulateInt32SliceRec: ", PopulateInt32SliceRec(10, []int32{}))
		return nil
	},
}

// Recursive functions in go

// Reverse a slice recursively putting the last element of the input slice
// and append it to a accumulator slice (starts empty)
func Reverse(x []int32, acc []int32) []int32 {
	// stop condition
	if len(x) == 0 {
		return acc
	} else {
		// len(x)-1 = last element
		t := x[len(x)-1]
		acc = append(acc, t)
		// x[:len(x)-1] = create a subslice that DO NOT contains the last element
		return Reverse(x[:len(x)-1], acc)
	}
}

// Recursively traverse a slice and mount it in direct order by removing the first element
// of the input slice and append it to a accumulator slice (starts empty)
func Direct(x []int32, acc []int32) []int32 {
	// stop condition
	if len(x) == 0 {
		return acc
	} else {
		t := x[0]
		acc = append(acc, 2*t)
		return Direct(x[1:len(x)], acc)
	}
}

// Apply the function recursively with an operator between a numer and its predecessor
// earlier function version tests if x equals 1 as stop condition
// This is a security risk as it allows one to pass a negative number
// which throws stack overflow error
func Operator(x int32) int32 {
	//stop condition
	if x <= 1 {
		return 1
	} else {
		return x + Operator(x-1)
	}

}

func Operator2(x []int32) int32 {
	//stop condition
	if len(x) == 1 {
		return x[0]
	} else {
		return x[0] + Operator2(x[1:])
	}

}

func PopulateInt32Slice(n int32) []int32 {
	result := make([]int32, n)
	for i := int32(0); i < n; i++ {
		result[i] = i + 1
	}
	return result
}

// recursive version
func PopulateInt32SliceRec(n int32, acc []int32) []int32 {
	if int(n)-len(acc) == 0 {
		return acc
	}
	acc = append(acc, int32(len(acc))+1)
	return PopulateInt32SliceRec(n, acc)
}

func RecSort(li []int) []int {
	if len(li) == 0 {
		return []int{}
	}
	
	var gtr, io int
	for i, v := range li {
		if v > gtr {
			gtr = v
			io = i
		}
	}
	return append(sort(append(li[:io], li[io+1:]...)), gtr)
}

func RecCount(n int) []int {
	if n == 0 {
		return []int{}
	}
	return append([]int{n}, count(n-1)...)
}

