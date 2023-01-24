package lg_problems

import (
	"log"
	"reflect"
)

var MinimumSwaps2_arr = []int32{7, 1, 3, 2, 4, 5, 6}

func MinimumSwaps2(a []int32) int32 {
	b := initArray(len(a))
	var s int32
	for i, _ := range b {
		swap(b, a[int32(i)], b[int32(i)])
		s = s + 1

		log.Print(b)
		if reflect.DeepEqual(a, b) {
			break
		}
	}
	log.Print(s - 1)
	return s - 1

}

func initArray(l int) []int32 {
	a := make([]int32, l)
	for i := int32(0); i < int32(l); i++ {
		a[i] = i + 1
	}
	return a
}

func swap(a []int32, x int32, y int32) {
	var ix, iy int32
	for i, v := range a {
		if v == x {
			ix = int32(i)
		}
		if v == y {
			iy = int32(i)
		}
	}
	a[ix], a[iy] = a[iy], a[ix]
}

/* This code was kept here only to document the algorithm design process

func MinimumSwaps2(a []int32) int32 {
	b := initArray(len(a))
	var s int32
	for !reflect.DeepEqual(a, b) {
		for i, _ := range b {
			swap(b, int32(i), a[int32(i)]-1)
			s = s + 1
			log.Print(b)
		}
		if reflect.DeepEqual(a, b) {
			return s
		}
	}
	return s
}


func swap(a []int32, x int32, y int32) {
	a[x], a[y] = a[y], a[x]
}

************************************************************************/

/* This code was kept here only to document the algorithm design process

func MinimumSwaps2(a []int32) int32 {
	// Given a number that isn't in its place
	// Test its displacement with every other number
	// displacement = proper position - actual position
	// displacement(7) => 6 - 0 = 6
	// Just let my instinct
	// It is hard to see a stop condition

	for {

		var x, y int32
		for i := int32(0); i < int32(len(a)); i++ {

			if a[i]-1-i > 0 {
				if a[i]-1-i > (a[x])-1-x {
					x = i
				}
			} else {
				if a[i]-1-i < (a[y])-1-y {
					y = i
				}
			}
			log.Print(displacement(int32(i), a[i]), x, y)

		}
		swap(a, x, y)
		log.Print(a)
	}
	return a[0]
}

func displacement(i int32, v int32) int32 {
	return (v - 1) - i
}

*****************************************************************/
