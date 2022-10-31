package lg_misc

func SelectionSort(a []int32) []int32 {
	var m int32
	// traverse the array reward (from last to first)
	// find the greatest and put it at the end of the unsorted array

	// i equals to the position of the last element of unsorted array
	for i := int32(len(a) - 1); i > 0; i-- {
		m = i
		for j := int32(0); j < i; j++ {
			// if some element of the unsorted array is greater
			// than the
			if a[m] < a[j] {
				// position of greatest element temporarily
				// stored in m
				m = j
			}
			// put last element of unsorted array in the place of greatest element
			// put the greatest element as last element
			a[m], a[i] = a[i], a[m]

		}

	}
	return a
}

func InsertionSort(a []int32) []int32 {
	var m int32
	for i, _ := range a[1:] {
		j := i
		for j > 0 && a[j] < a[j-1] {
			a[j-1], a[j] = a[j], a[j-1]
			j = j - 1
		}
	}
}
