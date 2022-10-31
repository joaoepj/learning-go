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
		acc = append(acc, t)
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

// Apply the function recursively as an operator between list elements
func Operator(x int32) int32 {
	//stop condition
	if x == 1 {
		return 1
	} else {
		return x + Operator(x-1)
	}

}

func Operator2(x []int32) int32 {
	//stop condition
	if len(x) == 1 {
		return x[0]
	} else {
		return x[0] + Operator2(x[1:len(x)])
	}

}

func PopulateInt32Slice(n int32) []int32 {
	result := make([]int32, n)
	for i := int32(0); i < n; i++ {
		result[i] = i + 1
	}
	return result
}
