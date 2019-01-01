package main

import "testing"

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(40)
	}
}
