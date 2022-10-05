package lg_problems

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var p, n, z int32 = 0, 0, 0
	for _, v := range arr {
		switch {
		case v > 0:
			p++
		case v < 0:
			n++
		case v == 0:
			z++
		}
	}
	l := len(arr)
	pr := float64(p) / float64(l)
	nr := float64(n) / float64(l)
	zr := float64(z) / float64(l)

	fmt.Printf("%.6f\n", pr)
	fmt.Printf("%.6f\n", nr)
	fmt.Printf("%.6f\n", zr)

}

func constrain1(n int) bool {
	return 0 < n && n <= 100
}

func constrain2(arr []int32) bool {
	res := true
	for _, v := range arr {
		if -100 > v && v > 100 {
			res = false
			break
		}
	}
	return res
}
