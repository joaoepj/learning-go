package lg_algorithms

import (
	//"reflect"
	"testing"
)

func BenchmarkRecFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecFibonacci(50)
	}
}

func BenchmarkCacheFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CacheFibonacci(cfMap, 50)
	}
}
