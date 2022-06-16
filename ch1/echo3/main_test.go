package main

import (
	"testing"
)

func BenchmarkTest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
