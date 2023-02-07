package lg_problems

import (
	"fmt"
	"log"
)

var (
	Magazine = []string{"two", "times", "three", "is", "not", "four"}
	Note     = []string{"two", "times", "two", "is", "four"}
)

func CheckMagazine(magazine []string, note []string) {
	// Write your code here
	magazineWC := make(map[string]int)
	noteWC := make(map[string]int)

	for _, word := range magazine {

		magazineWC[word] = magazineWC[word] + 1
	}
	log.Print(magazineWC)

	for _, word := range note {

		noteWC[word] = noteWC[word] + 1

	}
	log.Print(noteWC)
	for word, count := range noteWC {

		if !(magazineWC[word] >= count) {
			fmt.Print("No")
			return
		}
	}
	fmt.Print("Yes")
}
