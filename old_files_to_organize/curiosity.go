package main

import (
	"fmt"
)

const (
	one = 1
	two = 2
)

// Just curious this multiline variables declaration I saw in some code
var (
	a int
	b string
	c float64
)

func main() {
	a := 1
	b := "foo"
	c := 0.5
	fmt.Println(one+a, b, c)

}
