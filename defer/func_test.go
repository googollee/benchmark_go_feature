package deferbench

import "testing"

func BenchmarkDefer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useDefer(1)
	}
}

func BenchmarkNormal1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useNormal(1)
	}
}

func BenchmarkDefer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useDefer(4)
	}
}

func BenchmarkNormal4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useNormal(4)
	}
}
