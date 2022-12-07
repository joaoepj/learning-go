package lg_misc

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

// instead of printing this function should return some value that allow tests
// bus as the the tuple involve 3 diffrent types maybe we need a new type
func Curiosity() {
	a := 1
	b := "foo"
	c := 0.5
	fmt.Println(one+a, b, c)

}
