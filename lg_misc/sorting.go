package lg_misc

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var SCSorting *cli.Command = &cli.Command{
	Name:    "sorting",
	Aliases: []string{"st"},
	Usage:   "Sorting Functions",
	Action: func(c *cli.Context) error {
		randomArray := RandomIntegers32(10)
		fmt.Println("PrintArray: ")
		PrintArray(randomArray)
		// Destroying randomArray
		//fmt.Println("SelectionSort2: ", SelectionSort2(randomArray))
		fmt.Println("SelectionSort: ", SelectionSort(randomArray))
		fmt.Println("InsertionSort: ", InsertionSort(randomArray))
		fmt.Println("SortQueries:", SortQueries(Queries))
		return nil
	},
}

var Queries = [][]int32{
	{1, 5, 3},
	{4, 8, 7},
	{6, 9, 1},
}

func SelectionSort2(a []int32) []int32 {
	var minor int32

	for i := int32(0); i < int32(len(a)); i++ {
		minor = a[i]
		for j := i + 1; j < int32(len(a)); j++ {
			if a[j] < minor {
				minor = a[j]
				fmt.Println(minor)
			}
			fmt.Println(a)
		}
		a[i] = minor
	}
	return a
}

func PrintArray(a []int32) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func SelectionSort(a []int32) []int32 {
	var m int32
	// traverse the array reward (from last to first)
	// find the greatest element and put it at the end of the unsorted array

	// i equals to the position of the last element of unsorted array
	for i := int32(len(a) - 1); i > 0; i-- {
		m = i
		for j := int32(0); j <= i; j++ {
			// if some element of the unsorted array is greater
			// than the the first element of sorted array
			if a[m] < a[j] {
				// position of greatest element temporarily
				// stored in m
				m = j

			}
			// put last element of unsorted array in the place of greatest element
			// put the greatest element as last element
			a[m], a[i] = a[i], a[m]

		}

	}
	return a
}

func InsertionSort(a []int32) []int32 {
	for i, _ := range a[0:] {
		j := i
		for j > 0 && a[j] < a[j-1] {
			a[j-1], a[j] = a[j], a[j-1]
			j = j - 1
		}
	}
	return a
}

func SortQueries(q [][]int32) [][]int32 {
	if len(q) == 0 {
		return [][]int32{}
	}

	tmp_row := make([]int32, 3)
	//tmp_matrix := make([][]int32, 0)
	var oi int32
	for i, row := range q {
		// get the greater weight
		if row[2] > tmp_row[2] {
			tmp_row = row
			oi = int32(i)

		}
	}
	return append(q[:oi], q[oi+1:]...)
}
