package lg_problems

import (
	"fmt"
)

var TSMajor, TSMinor string = "abxdefghijabcofjdedlnkdflxkvlabc", "abc"

/*
Time spent reading, thinking and typing:
 17:37:4800:12:53.4292023-02-15

Result:
Time limit exceeded
Your code did not execute within the time limits. Please optimize your code. For more information on execution time limits, refer to the environment page

Environment constrains:
LANGUAGE	VERSION	LIMITS	LIBRARIES
TIME (SECS)	MEMORY (MB)

Go			1.13	4 secs	1024 MB
encoding/json,
encoding/csv,
encoding/xml,
strings, math,
container/heap,
container/list

*/

func TwoStrings(s1 string, s2 string) string {
	// Write your code here
	var match string = "NO"
	var minor, major string

	if len(s1) >= len(s2) {
		minor = s2
		major = s1
	} else {
		minor = s1
		major = s2
	}

	//O(n^2)
	for i := 0; i < len(minor); i++ {
		for _, v := range major {
			if minor[i] == byte(v) {
				match = "YES"
			}
		}
	}

	return match

}

func TwoStrings1(s1 string, s2 string) {
	var major, minor string
	if len(s1) >= len(s2) {
		major = s1
		minor = s2
	} else {
		major = s2
		minor = s1
	}
	found := false
	for i, letterM := range major {
		if rune(minor[0]) == letterM {
			for j := 1; j < len(minor); j++ {
				if minor[j] != major[i+j] {
					break
				}
				if j == len(minor)-1 {
					found = true
					fmt.Print("YES")
					return
				}
			}

		}

	}
	if !found {
		fmt.Print("NO")
	}
}
