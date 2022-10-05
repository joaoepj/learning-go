package lg_problems

import "os"

// A Solution for HackerHank simpleArraySum Problem
func simpleArraySum(ar []int32) int32 {
	// Write your code here
	if !constrain_sas(int32(len(ar))) {
		os.Exit(1)
	}

	for i := 0; i < len(ar); i++ {
		if !constrain_sas(ar[i]) {
			os.Exit(1)
		}
	}

	var acc int32 = 0
	for i := 0; i < len(ar); i++ {
		acc = acc + ar[i]
	}

	return acc
}

func constrain_sas(x int32) bool {
	return x > 0 && x <= 1000
}
