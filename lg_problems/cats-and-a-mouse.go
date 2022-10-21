package lg_problems

func myAbs(x int32) int32 {
	if x < 0 {
		// Ops, just the minus would suffice
		//return int32(-1 * x)
		return -x
	} else {
		return x
	}
}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	var result string
	if myAbs(x-z) == myAbs(z-y) {
		result = "Mouse C"
	}
	if myAbs(x-z) < myAbs(z-y) {
		result = "Cat A"
	}
	if myAbs(x-z) > myAbs(z-y) {
		result = "Cat B"
	}
	return result
}
