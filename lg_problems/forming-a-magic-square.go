package lg_problems

import "fmt"

type (
	// pair of pairs
	Pop struct {
		u, v, x, z int32
	}
)

// return difference to five
func five_t(x int32) int32 {
	if x > 5 {
		return x - 5
	} else {
		return 5 - x
	}
}

// return the balancing value
func five_bal(x int32) int32 {
	return -x + 5
}

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	var result int32
	var t = []Pop{{0, 0, 2, 2}, {0, 1, 2, 1}, {0, 2, 2, 0}, {1, 2, 1, 0}}
	for _, v := range t {
		if five_t(s[v.u][v.v]) != five_t(s[v.x][v.z]) {
			fmt.Printf("Ops %d %d ", s[v.u][v.v], s[v.x][v.z])
			if bal := five_t(s[v.u][v.v]); bal != 0 {
				fmt.Printf("Bal %d ", bal)
				s[v.x][v.z] = 5 - bal
				fmt.Printf("Pos %d %d ", s[v.u][v.v], s[v.x][v.z])
				result = result + bal
			}
		}
	}
	return result
}
