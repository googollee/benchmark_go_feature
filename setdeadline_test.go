package setdeadline

import "testing"

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

func BenchmarkNotTimeout(b *testing.B) {
	svr, err := NewEcho("localhost:0")
	if err != nil {
		b.Fatal("listen error:", err)
	}
	addr := svr.Addr().String()
	cnts, err := DialMutiple(addr, SetDeadLineNon, 100)
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
