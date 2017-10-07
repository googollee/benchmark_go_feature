package setdeadline

import (
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func BenchmarkSetDeadLineEach(b *testing.B) {
	svr, err := NewEcho("localhost:0")
	if err != nil {
		b.Fatal("listen error:", err)
	}
	addr := svr.Addr().String()
	cnts, err := DialMutiple(addr, SetDeadLineEach, 100)
	if err != nil {
		b.Fatal("dial error:", err)
	}
	b.ResetTimer()

	cnts.Do(b.N)

	b.StopTimer()

	cnts.Close()
	svr.Close()
	svr.Wait()
	if err := cnts.Error(); err != nil {
		b.Fatal("client error:", err)
	}
}

func BenchmarkSetDeadLineBoth(b *testing.B) {
	svr, err := NewEcho("localhost:0")
	if err != nil {
		b.Fatal("listen error:", err)
	}
	addr := svr.Addr().String()
	cnts, err := DialMutiple(addr, SetDeadLineBoth, 100)
	if err != nil {
		b.Fatal("dial error:", err)
	}
	b.ResetTimer()

	cnts.Do(b.N)

	b.StopTimer()

	cnts.Close()
	svr.Close()
	svr.Wait()
	if err := cnts.Error(); err != nil {
		b.Fatal("client error:", err)
	}
}

func BenchmarkTimeoutLive(b *testing.B) {
	svr, err := NewEcho("localhost:0")
	if err != nil {
		b.Fatal("listen error:", err)
	}
	addr := svr.Addr().String()
	cnts, err := DialMutiple(addr, SetDeadLineLive, 100)
	if err != nil {
		b.Fatal("dial error:", err)
	}
	b.ResetTimer()

	cnts.Do(b.N)

	b.StopTimer()

	cnts.Close()
	svr.Close()
	svr.Wait()
	if err := cnts.Error(); err != nil {
		b.Fatal("client error:", err)
	}
}

func BenchmarkTimeoutNone(b *testing.B) {
	svr, err := NewEcho("localhost:0")
	if err != nil {
		b.Fatal("listen error:", err)
	}
	addr := svr.Addr().String()
	cnts, err := DialMutiple(addr, SetDeadLineNone, 100)
	if err != nil {
		b.Fatal("dial error:", err)
	}
	b.ResetTimer()

	cnts.Do(b.N)

	b.StopTimer()

	cnts.Close()
	svr.Close()
	svr.Wait()
	if err := cnts.Error(); err != nil {
		b.Fatal("client error:", err)
	}
}
