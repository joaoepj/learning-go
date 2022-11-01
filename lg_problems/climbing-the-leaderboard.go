package lg_problems

import (
	"fmt"

	lg_misc "github.com/joaoepj/learning-go/lg_misc"
)

// It works! But not in time
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here

	var joined, result []int32

	joined = ranked
	for w := int32(0); w < int32(len(player)); w++ {

		joined = append(joined, player[w])
		//fmt.Println("joined: ",joined)
		sorted := lg_misc.InsertionSort(joined)
		//fmt.Println("sorted: ",sorted)

		m := make(map[int32]int32)
		var rank, score int32 = 0, 1<<31 - 1
		for i := int32(len(sorted) - 1); i >= 0; i-- {
			if sorted[i] < score {
				score = sorted[i]
				rank++
				m[score] = rank
			}
		}

		k := player[w]
		if v, found := m[k]; found {
			result = append(result, v)
		}

		fmt.Println("player: ", player[w], "result: ", result, "m: ", m)
	}
	return result
}
