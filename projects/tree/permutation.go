package main

import "fmt"

const DEBUG = true

type Tree map[int]Tree

func exclude(i int, li []int) []int {
	var result []int
	for _, v := range li {
		if v != i {
			result = append(result, v)
		}
	}
	return result
}

func log(itf ...interface{}) {
	if DEBUG {
		fmt.Println(itf...)
	}
}

func permutation(t *Tree, li []int) [][]int {
	var loLists [][]int

	for _, v := range li {
		(*t)[v] = make(Tree)
		log("t:", t, "v:", v)
		c := (*t)[v]
		perm := permutation(&c, exclude(v, li))
		log("perm", perm)
		loLists = append(loLists, perm...)
	}
	return loLists
}

func main() {
	t := make(Tree)
	log("exclude", exclude(2, []int{1, 2, 3}))
	fmt.Println(permutation(&t, []int{1, 2, 3}))

}
