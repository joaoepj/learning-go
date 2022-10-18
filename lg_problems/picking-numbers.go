package lg_problems

var PickingNumbers_arr = []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}

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
