package reflectcall

import (
	"context"
	"reflect"
	"testing"
)

func BenchmarkNormalCall(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ret, err := F(ctx, "abc")
		_ = ret
		_ = err
	}
}

func BenchmarkReflectCall(b *testing.B) {
	ctx := context.Background()
	fv := reflect.ValueOf(F)
	ctxv := reflect.ValueOf(ctx)
	sv := reflect.ValueOf("abc")
	in := []reflect.Value{ctxv, sv}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out := fv.Call(in)
		_ = out[0].Interface().(string)
		if !out[1].IsNil() {
			_ = out[1].Interface().(error)
		}
	}
}
