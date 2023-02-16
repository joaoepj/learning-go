package lg_problems

import (
	"fmt"
	"log"
)

var ArrayManipulation_arr = [][][]int32{
	{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	},
	{

		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	},
	{
		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{5, 9, 15},
	},
}

type tuple struct {
	set byte
	elm int32
}

func ArrayManipulation(n int32, queries [][]int32) int64 {
	var result int64
	var min, max int32 = 1<<31 - 1, 0
	baseWeights := make(map[int64]map[int32]struct{})

	for _, set := range queries {
		if set[0] < min {
			min = set[0]
		}
		if set[1] > max {
			max = set[1]
		}

		mappedElement := make(map[int32]struct{})
		for element := set[0]; element <= set[1]; element++ {
			if _, ok := baseWeights[int64(set[2])]; ok {
				mappedElement = baseWeights[int64(set[2])]
			}
			mappedElement[element] = struct{}{}
			baseWeights[int64(set[2])] = mappedElement
			fmt.Println(baseWeights)
		}

	}
	fmt.Println(baseWeights)

	summedWeights := make([]int64, 0)
	for element := min; element <= max; element++ {
		var summedWeight int64 = 0
		for baseWeight, mappedElements := range baseWeights {

			if _, ok := mappedElements[element]; ok {
				_ = baseWeight
				summedWeight += baseWeight
			}
		}
		summedWeights = append(summedWeights, summedWeight)
	}

	fmt.Println(summedWeights)

	for _, weight := range summedWeights {
		if weight > result {
			result = weight
		}
	}
	fmt.Println(result)
	return result
}

/*
set		elements				weight
a		1 2 3 4 5				3
b		  2 3 4 5				7
c		    3 4					1


map[1:map[1:{} 2:{} 3:{} 4:{} 5:{} 6:{} 7:{} 8:{}] 7:map[3:{} 4:{} 5:{}] 8:map[2:{} 3:{} 4:{} 5:{} 6:{}] 15:map[5:{} 6:{} 7:{} 8:{} 9:{}]]
[0 1 9 16 16 31 24 16 16 15]
31
*/

func ArrayManipulation3(n int32, queries [][]int32) int64 {
	var result int64
	var min, max int32 = 1<<31 - 1, 0
	baseWeights := make(map[int64]map[int32]struct{})

	for _, set := range queries {
		if set[0] < min {
			min = set[0]
		}
		if set[1] > max {
			max = set[1]
		}

		mappedElement := make(map[int32]struct{})
		for element := set[0]; element <= set[1]; element++ {
			if _, ok := baseWeights[int64(set[2])]; ok {
				mappedElement = baseWeights[int64(set[2])]
			}
			mappedElement[element] = struct{}{}
			baseWeights[int64(set[2])] = mappedElement
			fmt.Println(baseWeights)
		}

	}
	fmt.Println(baseWeights)

	summedWeights := make([]int64, 0)
	for element := min; element <= max; element++ {
		var summedWeight int64 = 0
		for baseWeight, mappedElements := range baseWeights {

			if _, ok := mappedElements[element]; ok {
				_ = baseWeight
				summedWeight += baseWeight
			}
		}
		summedWeights = append(summedWeights, summedWeight)
	}

	fmt.Println(summedWeights)

	for _, weight := range summedWeights {
		if weight > result {
			result = weight
		}
	}
	fmt.Println(result)
	return result
}

/*
map[8:[{0 2} {0 3} {0 4} {0 5} {0 6}]]
map[8:[{0 2} {0 6} {0 4} {0 5}] 15:[{0 3} {1 3}]]
map[8:[{0 2} {0 6} {0 5}] 15:[{0 3} {1 3} {0 4} {1 4}]]
map[8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2} {2 3}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2} {2 3} {2 4}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2} {2 3} {2 4} {2 5}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[1:[{3 5}] 7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[1:[{3 5} {3 6}] 7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[1:[{3 5} {3 6} {3 7}] 7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[1:[{3 5} {3 6} {3 7} {3 8}] 7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[1:[{3 5} {3 6} {3 7} {3 8} {3 9}] 7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
map[1:[{3 5} {3 6} {3 7} {3 8} {3 9}] 7:[{2 1} {2 2} {2 3} {2 4} {2 5} {2 6} {2 7} {2 8}] 8:[{0 2} {0 6}] 15:[{0 3} {1 3} {0 4} {1 4} {0 5} {1 5}]]
15

*/

/*
set		elements				weight
a		1 2 3 4 5				3
b		      4 5 6 7 8			7
c		          6 7 8 9		1

map[3:(a,1),(a,2),(a,3),(a,4),(a,5)]

map[3:(a,1),(a,2),(a,3),(a,5),10:[[(a,4),(b,4)],[(a,5),(b,5)]]

for each set

	for each element
		if is the first set
			if map is empty
				add current weight as key
			if tuple (set, element) is not in value
				search for tuple in weight's value
		if is second set onwards
			if current element is present in a tuple of previous value
				remove previous tuple and append to current value
				append current tuple to current value
				current key = current weight + previous weight
				add current key and current value to map
			if current element is not present in a tuple of previous value
				append current tuple to current value
				add current key and current value to map
*/

/* This code was kept here only to document the algorithm design process */
func ArrayManipulation2(n int32, queries [][]int32) int64 {
	var result int64
	weightSums := make(map[int64][]tuple)
	for s, set := range queries {
		if s == 0 {
			if _, ok := weightSums[int64(set[2])]; !ok {
				tuples := []tuple{}
				for element := set[0]; element <= set[1]; element++ {
					tuples = append(tuples, tuple{set: byte(s), elm: element})
				}
				weightSums[int64(set[2])] = tuples
				fmt.Println(weightSums)
				continue
			}

		}
		tuples := []tuple{}
		for element := set[0]; element <= set[1]; element++ {
			var removed bool
			for w := range weightSums {
				removed, weightSums[w] = removeTuple(element, weightSums[w])
				if removed {
					tuples = append(tuples, tuple{set: byte(s - 1), elm: element}, tuple{set: byte(s), elm: element})
					weightSum := int64(w + int64(queries[s][2]))
					weightSums[weightSum] = tuples
					break
				} else {
					tuples = append(tuples, tuple{set: byte(s), elm: element})
					weightSums[int64(queries[s][2])] = tuples
					break
				}
			}
			fmt.Println(weightSums)
		}
	}
	fmt.Println(weightSums)
	for key, _ := range weightSums {
		if key > result {
			result = key
		}
	}
	fmt.Println(result)
	return int64(0)

}

func removeTuple(elm int32, lt []tuple) (bool, []tuple) {
	ok := false
	for i, wanted := range lt {
		if elm == wanted.elm {
			lt[i] = lt[len(lt)-1]
			lt = lt[:len(lt)-1]
			ok = true
		}
	}
	return ok, lt
}

/* This code was kept here only to document the algorithm design process */

// set intersection
func ArrayManipulation1(n int32, queries [][]int32) int64 {
	var result int64
	intersection := make(map[int32]int64)
	for _, row := range queries {
		for i := row[0]; i <= row[1]; i++ {
			intersection[i] = intersection[i] + int64(row[2])
		}
	}
	log.Print(intersection)
	for _, v := range intersection {
		if v > result {
			result = v
		}
	}
	return result
}
