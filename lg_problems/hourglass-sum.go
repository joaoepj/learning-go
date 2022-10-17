package lg_problems

import "fmt"

var HourglassSum_arr = [][]int32{
	{1, 1, 1, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 0},
	{0, 9, 2, -4, -4, 0},
	{0, 0, 0, -2, 0, 0},
	{0, 0, -1, -2, -4, 0},
}

func HourglassSum(arr [][]int32) int32 {
	// Write your code here
	var result, sum int32 = -63, -63
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			sum = arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			fmt.Printf("[%d,%d]=%d ", i, j, sum)
			if sum > result {
				result = sum
			}
		}
	}
	return result
}
