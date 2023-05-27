package main

import "fmt"

func count(n int) []int {
	if n == 1 { return []int{1}}

	return append(count(n-1), n)
}

func main() {
	fmt.Println(count(10))
}