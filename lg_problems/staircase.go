package problems

import (
	"fmt"
)

// HackerHank Staircase Problem

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func Staircase(n int32) {
	// Write your code here
	for line := 1; int32(line) <= n; line++ {
		stairline(int32(line), n)
	}

}

func stairline(a int32, b int32) {
	var chars []byte

	for i := 0; int32(i) < (b - a); i++ {
		chars = append(chars, ' ')
	}
	for i := 0; int32(i) < a; i++ {
		chars = append(chars, '#')
	}
	fmt.Println(string(chars))
}
