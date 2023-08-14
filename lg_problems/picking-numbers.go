package lg_problems

import (
	"fmt"

	"github.com/joaoepj/learning-go/lg_misc"
	"github.com/urfave/cli/v2"
)

var PickingNumbers_arr = [][]int32{
	{4, 6, 5, 3, 3, 1},
	{1, 1, 2, 2, 4, 4, 5, 5, 5},
}

// SC stands for Sub Command, refering to the sub command concept
// of urfave/cli/v2 library. It means that below/inside the command
// 'problem' there is a sub command pickingnumbers, i.e.
var SCPickingNumbers *cli.Command = &cli.Command{
	Name:    "pickingnumbers",
	Aliases: []string{"pn"},
	Usage:   "Picking Numbers Problem",
	Action: func(c *cli.Context) error {
		fmt.Println("PickingNumbers1")
		for _, test := range PickingNumbers_arr {
			PickingNumbers1(test)
		}

		fmt.Println("PickingNumbers2")
		for _, test := range PickingNumbers_arr {
			PickingNumbers2(test)
		}
		fmt.Println("PickingNumbers3")
		for _, test := range PickingNumbers_arr {
			PickingNumbers3(test)
		}

		return nil
	},
}

// third submission, a year later
func PickingNumbers3(a []int32) int32 {
	// Write your code here

	var subArrayMap map[int32][]int32
	subArrayMap = make(map[int32][]int32)
	// 1. you should traverse the 'a' array and
	// build the greatest array you can
	// 2. Return the lenght of builded array
	for i := 0; i < len(a); i++ {
		sb, ok := subArrayMap[a[i]]
		if !ok {
			var tmpsb []int32
			tmpsb = append(tmpsb, a[i])
			subArrayMap[a[i]] = tmpsb

			sb, _ = subArrayMap[a[i]]

			for j := i + 1; j < len(a); j++ {
				if a[j]-sb[len(sb)-1] <= 1 {
					// beware, updating retrieved values will not update the map
					// sb = append(sb, a[j])
					subArrayMap[a[i]] = append(subArrayMap[a[i]], a[j])
				}
			}

		}

		_ = sb

		fmt.Println(subArrayMap)
	}

	var saLenght int32
	for _, subArray := range subArrayMap {
		if len(subArray) > int(saLenght) {
			saLenght = int32(len(subArray))
		}
	}

	fmt.Println(saLenght)
	return saLenght

}

// second submission
func PickingNumbers2(a []int32) int32 {

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

// first submission
func PickingNumbers1(a []int32) int32 {
	var result int32
	//sb stands for subarrays
	// int is better for sorting
	var sb map[int]int
	// any operation over uninitialized map will lead to crash
	// uncomment the line below to sse it in action
	// sb[0] = 0
	sb = make(map[int]int)
	//sb[len(sb)] = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] <= 1 {
			sb[len(sb)]++
		} else {
			sb[len(sb)+1] = 0
		}
	}
	//sort.Ints(sb)
	fmt.Println(result)
	return result
}
