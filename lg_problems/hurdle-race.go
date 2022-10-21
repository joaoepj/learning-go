package lg_problems

func HurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var result int32
	for _, v := range height { //O(n)
		if v > k && v-k > result { //O(1)
			result = v - k //O(1)
		}
	}
	return result
}
