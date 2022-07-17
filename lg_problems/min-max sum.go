package problems

import (
	"fmt"
	"os"
)

// A Solution for HackerHank Min-Max Sum Problem
func MiniMaxSum(arr []int32) {
	// Write your code here
	if !constrain1(arr) {
		os.Exit(1)
	}
	var tmp, min, max int = 0, 9223372036854775807, -9223372036854775808
	for i := 0; i < len(arr); i++ {
		tmp = sumArray(excludeElement(int32(i), arr))
		//fmt.Print(tmp)
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
	}
	fmt.Print(min, " ", max)
}

func constrain1(slc []int32) bool {
	result := false
	for i := 0; i < len(slc); i++ {
		if 1 <= slc[int32(i)] && slc[int32(i)] <= 1000000000 {
			result = true
		}
	}
	return result
}

// sumArray
func sumArray(slc []int32) int {
	var tmp int
	for _, v := range slc {
		tmp += int(v)
	}
	return tmp
}

// excludeElement
func excludeElement(e int32, slc []int32) []int32 {
	var tmp []int32
	for i := 0; int32(i) < e; i++ {
		tmp = append(tmp, slc[int32(i)])
	}
	for i := e + 1; int32(i) < int32(len(slc)); i++ {
		tmp = append(tmp, slc[int32(i)])
	}
	return tmp
}
