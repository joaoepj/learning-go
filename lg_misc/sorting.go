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
		randomArray := RandomIntegers32(5)
		fmt.Println("PrintArray: ")
		PrintArray(randomArray)
		// Destroying randomArray
		//fmt.Println("SelectionSort2: ", SelectionSort2(randomArray))
		fmt.Println("SelectionSort:", SelectionSort(randomArray))
		fmt.Println("InsertionSort:", InsertionSort(randomArray))
		fmt.Println("MergeSort:", MergeSort(randomArray, 0, len(randomArray)))
		fmt.Println("mergeSort:", mergeSort(randomArray))
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

func MergeSort(a []int32, start int, end int) []int32 {
	// stop condition
	if end-start > 1 {
		middle := (start + end) / 2
		L := a[:middle]
		R := a[middle:]
		fmt.Println(L, R)
		MergeSort(a, start, middle)
		MergeSort(a, middle, end)
		i, j := 0, 0
		for start < end {
			fmt.Println(a, L, R, start, middle, end, i, j)
			if j >= len(R) || (i < len(L) && L[i] < R[j]) {
				a[start] = L[i]
				i++
			} else {
				a[start] = R[j]
				j++
			}
			start++
		}

	}
	return a
}

// Seeing that my code would not work
// I decided to copy a working one
// https://blog.boot.dev/golang/merge-sort-golang/
// much simpler than mine, let's study it.
// P.S.: Needless to say that my code is already a copy from
// MIT course code. It is actually a transliteration
// from course python code to golang.
func mergeSort(items []int32) []int32 {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int32, b []int32) []int32 {
	final := []int32{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
