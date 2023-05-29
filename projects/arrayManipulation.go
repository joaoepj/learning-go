package main

import "fmt"

var arrayManipulationSample1 = [][]int32{
	{1, 5, 3},
	{4, 8, 7},
	{6, 9, 1},
}

var arrayManipulationSample2 = [][]int32{
	{1, 2, 100},
	{2, 5, 100},
	{3, 4, 100},
}

const DEBUG = true

func debug(a ...interface{}) {
	if DEBUG {
		fmt.Println(a...)
	}
}

func ArrayManipulation(n int32, queries [][]int32) int64 {
	acc := make([]int64, n)
	debug("1", acc)
	return recArrayManip(queries, acc)
}

func recArrayManip(queries [][]int32, acc []int64) int64 {
	if len(queries) == 1 {
		for i := queries[0][0]; i <= queries[0][1]; i++ {
			acc[i] += int64(queries[0][2])
		}
		debug("2", acc)
		sort(acc)
		debug("3", acc)
		return acc[len(acc)-1]
	}

	for i := queries[0][0]; i < queries[0][1]; i++ {
		acc[i] += int64(queries[0][2])
	}

	return recArrayManip(queries[1:], acc)

}

func sort(li []int64) []int64 {
	if len(li) == 1 {
		return li
	}

	var index int
	var greater int64
	for i, v := range li {
		if v > greater {
			greater = v
			index = i
		}
	}

	debug("index", index)
	debug("li[:index]", li[:index])
	debug("li[index+1:]", li[index+1:])
	return append(sort(append(li[:index], li[index+1:]...)), greater)
}

func main() {
	fmt.Println(ArrayManipulation(5, arrayManipulationSample2))
}
