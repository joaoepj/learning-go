package lg_misc

import (
	"math/rand"
	"testing"
)

// go test -run ^$ -bench ^BenchmarkRemoveVowels$ ./lg_misc/
func BenchmarkRemoveVowels(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withVowels := RandomString(rand.Intn(20000))
		RemoveVowels(withVowels)
	}

}

//func BenchmarkRemoveVowels2c(b *testing.B) { BenchmarkRemoveVowels(200, b) }
