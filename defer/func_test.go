package deferbench

import "testing"

func BenchmarkLoopDefer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useLoopDefer(1)
	}
}

func BenchmarkFuncDefer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useFuncDefer(1)
	}
}

func BenchmarkNormal1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useNormal(1)
	}
}

func BenchmarkLoopDefer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useLoopDefer(4)
	}
}

func BenchmarkFuncDefer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useFuncDefer(4)
	}
}

func BenchmarkNormal4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useNormal(4)
	}
}
