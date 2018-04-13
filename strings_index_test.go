package strings_index

import (
	"strings"
	"testing"
)

func BenchmarkIndex1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Index("abcabcabcdefghi", "e")
	}
}

func BenchmarkIndex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Index("abcabcabcdefghi", "z")
	}
}

func BenchmarkIndexByte1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.IndexByte("abcabcabcdefghi", 'e')
	}
}

func BenchmarkIndexByte2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.IndexByte("abcabcabcdefghi", 'z')
	}
}

// Results (using go1.10 darwin/amd64):
// BenchmarkIndex1-8               300000000                4.01 ns/op
// BenchmarkIndex2-8               300000000                4.39 ns/op
// BenchmarkIndexByte1-8           1000000000               2.23 ns/op
// BenchmarkIndexByte2-8           1000000000               2.25 ns/op
