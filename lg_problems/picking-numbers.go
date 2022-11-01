package lg_problems

import (
	"fmt"

	"github.com/joaoepj/learning-go/lg_misc"
)

var PickingNumbers_arr = []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}

// first submission
func PickingNumbers(a []int32) int32 {
	var result int32
	//sb stands for subarrays
	// int is better for sorting
	var sb map[int]int
	//sb[len(sb)] = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] <= 1 {
			sb[len(sb)]++
		} else {
			sb[len(sb)+1] = 0
		}
	}
	//sort.Ints(sb)
	return result
}

// second submission
func pickingNumbers(a []int32) int32 {

	var result, start, end int32
	var sbl []int32
	// Write your code here
	for {
		// stop condition 1
		if end == int32(len(a)-1) {
			sbl = append(sbl, end-start+1)
			break
		}
		// sub array finished
		if lg_misc.Abs32(a[end+1]-a[end]) > 1 {
			sbl = append(sbl, end-start+1)
			start = end + 1
			end = end + 1
		} else {
			// building subarray
			end = end + 1
		}
	}
	fmt.Println("sbl: ", sbl)
	// search longest array
	for _, v := range sbl {
		if result < v {
			result = v
		}
	}
	// return
	return result

}
