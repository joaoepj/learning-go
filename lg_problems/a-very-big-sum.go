package lg_problems

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var res int64
	//if !constrain1(len(ar)) {
	//	panic(-1)
	//}
	for _, v := range ar {
		//if !constrain2(v) {
		//	panic(-1)
		//} else {
		res += v
		//}
	}
	return res
}
