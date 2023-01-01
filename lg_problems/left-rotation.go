package lg_problems

func RotLeft(a []int32, d int32) []int32 {
	// Write your code here
	var tmp []int32
	for i := d; i > 0; i-- {
		tmp = make([]int32, 0)
		tmp = append(tmp, a[1:]...)
		tmp = append(tmp, a[0])
		a = tmp
	}
	return tmp
}
