```sh
go test -bench=. ./...
goos: linux
goarch: amd64
pkg: benchmark_go_feature/cast
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkCast-12       	714698967	         1.624 ns/op
BenchmarkNonCast-12    	916796696	         1.170 ns/op
PASS
ok  	benchmark_go_feature/cast	2.568s
goos: linux
goarch: amd64
pkg: benchmark_go_feature/defer
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkLoopDefer1-12    	66192694	        17.53 ns/op
BenchmarkFuncDefer1-12    	284976273	         4.054 ns/op
BenchmarkNormal1-12       	978095668	         1.075 ns/op
BenchmarkLoopDefer4-12    	22414300	        53.11 ns/op
BenchmarkFuncDefer4-12    	102236121	        11.47 ns/op
BenchmarkNormal4-12       	539346729	         1.992 ns/op
PASS
ok  	benchmark_go_feature/defer	9.733s
goos: linux
goarch: amd64
pkg: benchmark_go_feature/setdeadline
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkSetDeadLineEach-2   	    1465	    771752 ns/op
BenchmarkSetDeadLineBoth-2   	    1442	    752134 ns/op
BenchmarkTimeoutLive-2       	    1416	    779238 ns/op
BenchmarkTimeoutNone-2       	    1647	    784510 ns/op
PASS
ok  	benchmark_go_feature/setdeadline	5.097s
```
