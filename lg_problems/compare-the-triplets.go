package lg_problems

import "os"

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var acc = []int32{0, 0}
	//Is a constrained?
	for _, value := range a {
		if !constrain(value) {
			os.Exit(1)
		}
	}
	// Is b constrained?
	for _, value := range a {
		if !constrain(value) {
			os.Exit(1)
		}
	}
	// Are it equal in lenght?
	if len(a) != len(b) {
		os.Exit(1)
	}

	for index, value := range a {
		if value > b[index] {
			acc[0] = acc[0] + 1
		}

		if value < b[index] {
			acc[1] = acc[1] + 1
		}
	}

	return acc

}

func constrain(x int32) bool {
	return x >= 1 && x <= 100
}
