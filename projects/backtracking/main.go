package main

import "fmt"


func backtrack(s []int, k int, z int) {
	var candidates [10]int
	var ncandidates int

	if isSolution(s, k, z) {
		processSolution(s, k, 1	z)
	} else {
		k = k + 1
		constructCandidates(s, k, z, candidates, &ncandidates)
		for i := 0; i < ncandidates; i++ {
			s[k] = candidates[i]
			backtrack(s, k, z)
		}
	}
}

func constructCandidates(s []int, k int, z int, c [10]int, nc *int) {
	c[0] = 1
	c[1] = 0
	*nc = 2
}

func isSolution(s[]int, k int, n int) bool {
	return k == n
}

func processSolution(s []int, k int, z int) {
	fmt.Print("{")
	for i := 0; i < k; i++ {
		if s[i] == 1 { fmt.Print(i+1)}
	}
	fmt.Println("}")
}

/*
The function rank is defined as:
rank(p) = (p1 - 1) * (|p| - 1)! + rank(p2,...,p|p|)
As example:
rank({2,1,3}) = 1 * 2! + rank({1,2}) = 2 + 0 * 1! + rank({1}) = 2
So, lets try the following:
rank({1,2,3}) = 0 * 2! + Still I cant figure it out what comes here

Maybe, the inverse operation unrank can shed some light in here.
unrank(2,3) tells that the first element of permutation must be '2'.
(2 - 1) * (3 - 1)! = 2

*/


func rank(perm []int) int {
	var tmp []int
	if len(perm) == 1 {return 0}

	v1 := perm[0] - 1
	if v1 > 0 { tmp = append(tmp, v1)}

	v2 := factorial((len(perm) -1))
	if v2 > 0 { tmp = append(tmp, v2)}

	fmt.Println(v1,v2, tmp)
	return v1 * v2 + rank(tmp) 
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i-1)
}

func main() {
	fmt.Println("factorial",factorial(0))

	for /*k*/_, v := range [][]int{{1,2,3},{3,2,1},{2,3,1},{1,3,2},{3,1,2},{2,1,3}} {
		fmt.Println("rank", v, rank(v))
	} 
	
	// s := []int{1,1,1}
	// for i := 0; i < len(s); i++ {
	// 	backtrack(s, 0, i)
	// }
	
}