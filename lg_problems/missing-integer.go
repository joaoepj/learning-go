// From https://codility.com
package lg_problems

import "fmt"

// There is a little problem I didnt bother to spot
// Maybe you are the luck one

var MissingInteger_arr = [][]int{
	{1, 3, 6, 4, 1, 2},
	{1, 2, 3},
	{-1, -3},
	{1, 3},
	{0, 3},
	{-5, 3},
}

func MissingInteger(A []int) int {
	var result int = -1
	var tmp int
	m := make(map[int]int)
	for _, v := range A {
		if v > 0 {
			m[v] = v
		}
	}

	for i := 2; i <= len(m); i++ {
		// if 1 is not present, then it is the answer
		if _, ok := m[1]; ok == false {
			result = 1
			break
		}
		// if v is present, it can be the last
		// else missing integer was found
		if v, ok := m[i]; ok == true {
			tmp = v
		} else {
			result = i
			break
		}

	}
	// if missing integer was not found, then it is the successor of last
	if result == -1 {
		result = tmp + 1
	}
	return result

}

func TestMissingInteger() {
	var x = [][]int{
		{1, 3, 6, 4, 1, 2},
		{1, 2, 3},
		{-1, -3},
		{1, 3},
		{0, 3},
		{-5, 3},
	}
	for _, v := range x {
		fmt.Println(v, " = ", MissingInteger(v))
	}
}
