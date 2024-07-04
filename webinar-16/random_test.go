package main

import "testing"

func BenchmarkGenerateLotsOfRandomNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLotsOfRandomNumbers()
	}
}

func BenchmarkGenerateLotsOfRandomNumbersConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLotsOfRandomNumbersConcurrent()
	}
}

// 15217	     77517 ns/op	   81920 B/op
