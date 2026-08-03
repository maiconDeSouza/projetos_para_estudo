package stringutils

import "testing"

func BenchmarkConcatenarComBuilder(b *testing.B) {
	for b.Loop() {
		ConcatenarComBuilder(1000)
	}
}
