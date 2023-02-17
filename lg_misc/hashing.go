package lg_misc

import "fmt"

var StringList = []string{
	"abcdefghabcdefgh",
	"abcdefgh",
	"abcdefg",
	"abcdef",
	"123def",
	"!@#$%&%@%@&#&#&#&",
	"def",
	"123",
	"a",
	"b",
}

func Hash(s string) uint64 {
	var hash uint64
	for _, c := range s {
		hash ^= uint64(c)
		hash *= 17
	}

	fmt.Println(hash)
	return hash
}
