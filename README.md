```sh
$ go test ./... -bench .
goos: linux
goarch: amd64
pkg: benchmark_go_feature/cast
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkCast-12       	672592316	         1.581 ns/op
BenchmarkNonCast-12    	866675088	         1.263 ns/op
PASS
ok  	benchmark_go_feature/cast	2.515s
goos: linux
goarch: amd64
pkg: benchmark_go_feature/defer
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkLoopDefer1-12    	63559786	        18.26 ns/op
BenchmarkFuncDefer1-12    	258426038	         4.198 ns/op
BenchmarkNormal1-12       	1000000000	         0.9327 ns/op
BenchmarkLoopDefer4-12    	22685126	        54.55 ns/op
BenchmarkFuncDefer4-12    	99063734	        11.92 ns/op
BenchmarkNormal4-12       	581541480	         1.875 ns/op
PASS
ok  	benchmark_go_feature/defer	9.636s
goos: linux
goarch: amd64
pkg: benchmark_go_feature/reflectcall
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkNormalCall-12     	84479767	        13.81 ns/op
BenchmarkReflectCall-12    	  995415	      1160 ns/op
PASS
ok  	benchmark_go_feature/reflectcall	3.684s
goos: linux
goarch: amd64
pkg: benchmark_go_feature/setdeadline
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkSetDeadLineEach-2   	    1477	    823700 ns/op
BenchmarkSetDeadLineBoth-2   	    1531	    789365 ns/op
BenchmarkTimeoutLive-2       	    1725	    804610 ns/op
BenchmarkTimeoutNone-2       	    1470	    787093 ns/op
PASS
ok  	benchmark_go_feature/setdeadline	7.433s
```
