package cast

import (
	"testing"
)

func BenchmarkCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CastFunc(i)
	}
}

func BenchmarkNonCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NonCastFunc(i)
	}
}
