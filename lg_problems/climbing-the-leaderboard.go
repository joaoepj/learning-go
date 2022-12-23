package lg_problems

import (
	"fmt"

	lg_misc "github.com/joaoepj/learning-go/lg_misc"
)

// It works! But not in time
func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
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


func ClimbingLeaderboard2(ranked []int32, player []int32) []int32 {
	// Write your code here
	var res []int32

	for _, game := range player {
		ranking := make([]Player, 0)
		var position int32
		var tmp int32 = 1<<31 - 1

		//append
		jointScores := append(ranked, game)
		fmt.Println("jointScores: ", jointScores)
		// sort
		sortedScores := lg_misc.InsertionSort(jointScores)
		fmt.Println("sortedScores: ", sortedScores)

		//build ranking

		// sortedScores is in crescent order
		// so traverse the slice from end to begin
		for i := int32(len(sortedScores) - 1); i >= 0; i-- {
			if sortedScores[i] < tmp { //fail if score equals maximum int32
				tmp = sortedScores[i]
				position = position + 1
			}
			ranking = append(ranking, Player{position, sortedScores[i]})
		}

		fmt.Println("ranking: ", ranking)

		// build response

		for _, player := range ranking {
			// the first is never repeated
			if player.score == game && len(res) == 0 {
				res = append(res, player.rank)
			}
			
			// if isn't the last, can be inserted
			if player.score == game && player.rank != res[len(res)-1] {
				res = append(res, player.rank)
			}
		}

		fmt.Println("res: ", res)
		fmt.Println()

	}

	return res
}