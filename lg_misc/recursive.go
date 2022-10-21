package lg_misc

// Recursive functions in go

// Reverse a slice recursively putting the last element of the input slice
// and append it to a accumulator slice (starts empty)
func Reverse(x []int32, acc []int32) []int32 {
	// stop condition
	if len(x) == 0 {
		return acc
	} else {
		// len(x)-1 = last element
		t := x[len(x)-1]
		y = append(acc, t)
		// x[:len(x)-1] = create a subslice that DO NOT contains the last element
		return Reverse(x[:len(x)-1], acc)
	}
}

// Recursively traverse a slice and mount it in direct order by removing the first element
// of the input slice and append it to a accumulator slice (starts empty)
func Direct(x []int32, acc []int32) []int32 {
	// stop condition
	if len(x) == 0 {
		return acc
	} else {
		t := x[0]
		acc = append(acc, 2*t)
		return Direct(x[1:len(x)], acc)
	}
}
