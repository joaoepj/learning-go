package lg_problems

import (
	"log"
)

/*var ArrayManipulation_arr = [][]int32{
	{1, 5, 3},
	{4, 8, 7},
	{6, 9, 1},
}*/

var ArrayManipulation_arr = [][]int32{
	{2, 6, 8},
	{3, 5, 7},
	{1, 8, 1},
	{5, 9, 15},
}

type tuple struct {
	set byte
	elm int32
}

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
func ArrayManipulation(n int32, queries [][]int32) int64 {
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
				log.Print(weightSums)
				continue
			}

		}
		tuples := []tuple{}
		for element := set[0]; element <= set[1]; element++ {
			var removed bool
			removed, weightSums[int64(queries[s-1][2])] = removeTuple(tuple{set: byte(s - 1), elm: element}, weightSums[int64(queries[s-1][2])])
			if removed {
				tuples = append(tuples, tuple{set: byte(s - 1), elm: element}, tuple{set: byte(s), elm: element})
				weightSum := int64(queries[s-1][2] + queries[s][2])
				weightSums[weightSum] = tuples
			} else {
				tuples = append(tuples, tuple{set: byte(s), elm: element})
				weightSums[int64(queries[s-1][2])] = tuples
			}
			log.Print(weightSums)
		}
	}
	log.Print(weightSums)
	for key, _ := range weightSums {
		if key > result {
			result = key
		}
	}
	log.Print(result)
	return int64(0)

}

func removeTuple(t tuple, lt []tuple) (bool, []tuple) {
	ok := false
	for i, wanted := range lt {
		if t == wanted {
			lt[i] = lt[len(lt)-1]
			lt = lt[:len(lt)-1]
			ok = true
		}
	}
	return ok, lt
}

/* This code was kept here only to document the algorithm design process

// set intersection
func ArrayManipulation(n int32, queries [][]int32) int64 {
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
*/
