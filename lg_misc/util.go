package lg_misc

func Abs32(x int32) int32 {
	if x < 0 {
		// Ops, just the minus would suffice
		//return int32(-1 * x)
		return -x
	} else {
		return x
	}
}
