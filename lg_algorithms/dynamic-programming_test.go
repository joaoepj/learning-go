package lg_algorithms

import (
	"testing"
)

func BenchmarkRecFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecFibonacci(50)
	}
}

func BenchmarkUltimateFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UltimateFibonacci(50)
	}
}

func BenchmarkCacheFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CacheFibonacci(cfMap, 50)
	}
}

func BenchmarkDynProgFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynProgFibonacci(50)
	}
}
