package lg_problems

import "github.com/joaoepj/learning-go/lg_misc"

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	var result string
	if lg_misc.Abs32(x-z) == lg_misc.Abs32(z-y) {
		result = "Mouse C"
	}
	if lg_misc.Abs32(x-z) < lg_misc.Abs32(z-y) {
		result = "Cat A"
	}
	if lg_misc.Abs32(x-z) > lg_misc.Abs32(z-y) {
		result = "Cat B"
	}
	return result
}
